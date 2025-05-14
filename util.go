package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/theckman/httpforwarded"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// cache SQL/HTML templates so repeated filesystem reads are not required
var globalTemplates map[string](*template.Template) = make(map[string](*template.Template))
var globalTemplatesMutex = &sync.Mutex{}

// formatBaseURL takes a hostname (baseHost) and a base path
// and joins them.  Both are parsed as URLs (using net/url) and
// then joined to ensure a properly formed URL.
// net/url does not support parsing hostnames without a scheme
// (e.g. example.com is invalid; http://example.com is valid).
// serverURLHost ensures a scheme is added.
func formatBaseURL(baseHost string, basePath string) string {
	urlHost, err := url.Parse(baseHost)
	if err != nil {
		log.Fatal(err)
	}
	urlPath, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(urlHost.ResolveReference(urlPath).String(), "/")
}

// serverURLBase returns the base server URL
// that the client used to access this service.
// All pg_tileserv routes are mounted relative to
// this URL (including path, if specified by the
// BasePath config option)
func serverURLBase(r *http.Request) string {
	baseHost := serverURLHost(r)
	basePath := viper.GetString("BasePath")

	return formatBaseURL(baseHost, basePath)
}

// serverURLHost returns the host (and scheme)
// for this service.
// In the case of access via a proxy service, if
// the standard headers are set, we return that
// hostname. If necessary the automatic calculation
// can be over-ridden by setting the "UrlBase"
// configuration option.
func serverURLHost(r *http.Request) string {
	// Use configuration file settings if we have them
	configURL := viper.GetString("UrlBase")
	if configURL != "" {
		return configURL
	}

	// Preferred scheme
	ps := "http"
	if r.TLS != nil {
		ps = "https"
	}

	// Preferred host:port
	ph := strings.TrimRight(r.Host, "/")

	// Check for the IETF standard "Forwarded" header
	// for reverse proxy information
	xf := http.CanonicalHeaderKey("Forwarded")
	if f, ok := r.Header[xf]; ok {
		if fm, err := httpforwarded.Parse(f); err == nil {
			if len(fm["host"]) > 0 && len(fm["proto"]) > 0 {
				ph = fm["host"][0]
				ps = fm["proto"][0]
				return fmt.Sprintf("%v://%v", ps, ph)
			}
		}
	}

	// Check the X-Forwarded-Host and X-Forwarded-Proto
	// headers
	xfh := http.CanonicalHeaderKey("X-Forwarded-Host")
	if fh, ok := r.Header[xfh]; ok {
		ph = fh[0]
	}
	xfp := http.CanonicalHeaderKey("X-Forwarded-Proto")
	if fp, ok := r.Header[xfp]; ok {
		ps = fp[0]
	}

	return fmt.Sprintf("%v://%v", ps, ph)
}

func getSQLTemplate(name string, tmpl string) *template.Template {
	tp, ok := globalTemplates[name]
	if ok {
		return tp
	}
	t := template.New(name)
	tp, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	globalTemplatesMutex.Lock()
	globalTemplates[name] = tp
	globalTemplatesMutex.Unlock()
	return tp
}

func renderSQLTemplate(name string, tmpl string, data interface{}) (string, error) {
	var buf bytes.Buffer
	t := getSQLTemplate(name, tmpl)
	err := t.Execute(&buf, data)
	if err != nil {
		return string(buf.Bytes()), err
	}
	sql := string(buf.Bytes())
	log.Debug(sql)
	return sql, nil
}

/******************************************************************************/

func getServerBounds(sridPtr *int) (b *Bounds, e error) {
	var srid int
	if sridPtr == nil {
		srid = globalDefaultCoordinateSystem
	} else {
		srid = *sridPtr
	}

	bounds, ok := globalServerBounds[srid]
	if ok {
		return bounds, nil
	}

	var (
		xmin, ymin, xmax, ymax float64
	)

	if globalProjectionBoundsTableName != "" {
		ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("DbTimeout")*time.Second)
		defer cancel()
		var err error
		xmin, ymin, xmax, ymax, err = dBProjectionBoundsRequest(ctx, srid)
		if err != nil {
			return nil, fmt.Errorf("failed to get projection bounds: %w", err)
		}

		log.Infof(
			"Using CoordinateSystem %d with bounds from %s [%g, %g, %g, %g]",
			srid,
			globalProjectionBoundsTableName,
			xmin, ymin, xmax, ymax,
		)
	} else {
		xmin = viper.GetFloat64(fmt.Sprintf("CoordinateSystem.%d.Xmin", srid))
		ymin = viper.GetFloat64(fmt.Sprintf("CoordinateSystem.%d.Ymin", srid))
		xmax = viper.GetFloat64(fmt.Sprintf("CoordinateSystem.%d.Xmax", srid))
		ymax = viper.GetFloat64(fmt.Sprintf("CoordinateSystem.%d.Ymax", srid))

		log.Infof(
			"Using CoordinateSystem.SRID %d with bounds [%g, %g, %g, %g]",
			srid,
			xmin, ymin, xmax, ymax,
		)
	}

	width := xmax - xmin
	height := ymax - ymin
	size := math.Min(width, height)

	/* Not square enough to just adjust */
	if math.Abs(width-height) > 0.01*size {
		return nil, errors.New("CoordinateSystem bounds must be square")
	}

	cx := xmin + width/2
	cy := ymin + height/2

	/* Perfectly square bounds please */
	xmin = cx - size/2
	ymin = cy - size/2
	xmax = cx + size/2
	ymax = cy + size/2

	bounds = &Bounds{srid, xmin, ymin, xmax, ymax}
	globalServerBounds[srid] = bounds
	return bounds, nil
}

func getTTL() (ttl int) {
	if globalTimeToLive < 0 {
		globalTimeToLive = viper.GetInt("CacheTTL")
	}
	return globalTimeToLive
}

/*****************************************************************************/
//Prometheus metrics collection

var (
	tilesProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "pg_tileserv_tile_requests_total",
		Help: "The total number of tiles processed",
	},
		[]string{
			"layer",
			"status_code",
		})
	tilesDurationHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "pg_tileserv_tile_requests_duration",
			Help:    "Tile request processing duration distribution",
			Buckets: []float64{.1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{
			"layer",
		},
	)
)

// metricsResponseWriter helps capture the HTTP status code
// of responses.
// Credit to github.com/Boerworz for the sample code
type metricsResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// NewMetricsResponseWriter instantiates and returns a metricsResponseWriter
func NewMetricsResponseWriter(w http.ResponseWriter) *metricsResponseWriter {
	return &metricsResponseWriter{w, http.StatusOK}
}

func (mrw *metricsResponseWriter) WriteHeader(code int) {
	mrw.StatusCode = code
	mrw.ResponseWriter.WriteHeader(code)
}

// tileMetrics returns a middleware that collects metrics for tile set endpoints.
// If EnableMetrics = false, a blank middleware is returned. This is to avoid all the Prometheus
// metrics operations from occuring if metrics are disabled.
//
// Requests that return an HTTP status in the range 400-499 are considered bad
// requests and are not tracked. This includes layers that do not exist (404) and
// invalid tiles (400). Server errors (500) will still be tracked.
// 404 and 400 errors cannot be tracked because label values would no longer be
// constrained to valid layers.
func tileMetrics(h http.Handler) http.Handler {
	if viper.GetBool("EnableMetrics") {

		// log metrics URL at startup
		basePath := viper.GetString("BasePath")
		log.Infof("Prometheus metrics enabled at %s/metrics", formatBaseURL(fmt.Sprintf("http://%s:%d",
			viper.GetString("HttpHost"), viper.GetInt("HttpPort")), basePath))

		// create the handler that will track metrics for tile requests.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// start a timer for the duration histogram.
			start := time.Now()

			mrw := NewMetricsResponseWriter(w)

			// get path variables from the router to determine the layer name
			vars := mux.Vars(r)
			layer := vars["name"]

			// call the next handler
			h.ServeHTTP(mrw, r)

			// Do not track metrics for invalid user requests (4xx errors)
			if mrw.StatusCode/100 == 4 {
				return
			}

			// get the counter for this request and then increment it
			counter, err := tilesProcessed.GetMetricWith(
				map[string]string{
					"layer":       layer,
					"status_code": strconv.Itoa(mrw.StatusCode),
				},
			)
			if err != nil {
				log.Warn("Unable to get tilesProcessed Prometheus counter.")
				return
			}
			// get the histogram metric and make an observation of the
			// response time.
			histogram, err := tilesDurationHistogram.GetMetricWith(
				map[string]string{
					"layer": layer,
				},
			)
			if err != nil {
				log.Warn("Unable to get tilesDurationHistogram Prometheus histogram.")
				return
			}

			counter.Inc()

			duration := time.Since(start)
			histogram.Observe(duration.Seconds())
		})
	}
	// if metrics are disabled, return a handler without any of the
	// metric operations.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
