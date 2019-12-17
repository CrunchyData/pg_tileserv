package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/theckman/httpforwarded"
	"net/http"
	"strings"
	// log "github.com/sirupsen/logrus"
)

func serverURLBase(r *http.Request) string {

	// Use configuration file settings if we have them
	if viper.GetString("UrlBase") != "" {
		return viper.GetString("UrlBase")
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
