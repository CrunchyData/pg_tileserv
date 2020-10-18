package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/theckman/httpforwarded"
	// log "github.com/sirupsen/logrus"
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
