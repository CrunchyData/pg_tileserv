package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/theckman/httpforwarded"
)

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
	configUrl := viper.GetString("UrlBase")
	if configUrl != "" {
		return configUrl
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
			ph = fm["host"][0]
			ps = fm["proto"][0]
			return fmt.Sprintf("%v://%v", ps, ph)
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

var globalTemplates map[string](*template.Template) = make(map[string](*template.Template))

func getSqlTemplate(name string, tmpl string) *template.Template {
	tp, ok := globalTemplates[name]
	if ok {
		return tp
	}
	t := template.New(name)
	tp, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	globalTemplates[name] = tp
	return tp
}

func renderSqlTemplate(name string, tmpl string, data interface{}) (string, error) {
	var buf bytes.Buffer
	t := getSqlTemplate(name, tmpl)
	err := t.Execute(&buf, data)
	if err != nil {
		return string(buf.Bytes()), err
	}
	sql := string(buf.Bytes())
	log.Debug(sql)
	return sql, nil
}

/******************************************************************************/

func getServerBounds() (b *Bounds, e error) {

	if globalServerBounds != nil {
		return globalServerBounds, nil
	}

	srid := viper.GetInt("CoordinateSystem.SRID")
	xmin := viper.GetFloat64("CoordinateSystem.Xmin")
	ymin := viper.GetFloat64("CoordinateSystem.Ymin")
	xmax := viper.GetFloat64("CoordinateSystem.Xmax")
	ymax := viper.GetFloat64("CoordinateSystem.Ymax")

	log.Infof("Using CoordinateSystem.SRID %d with bounds [%g, %g, %g, %g]",
		srid, xmin, ymin, xmax, ymax)

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

	globalServerBounds = &Bounds{srid, xmin, ymin, xmax, ymax}
	return globalServerBounds, nil
}
