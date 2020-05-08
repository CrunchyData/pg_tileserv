package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/theckman/httpforwarded"
	// log "github.com/sirupsen/logrus"
)

// serverURLBase returns the server URL
// that the client used to access this service.
// In the case of access via a proxy service, if
// the standard headers are set, we return that
// URL base. If necessary the automatic calculation
// can be over-ridden by setting the "UrlBase"
// configuration option
func serverURLBase(r *http.Request) string {
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

// Structure for filter row in postgres (incomes via http api)
type FilterData struct {
	FieldName string `json:"fieldName"`
	FieldType int    `json:"fieldType"`
	Operator  int    `json:"operator"`
	Arg0      string `json:"arg0,omitempty"`
	Arg1      string `json:"arg1,omitempty"`
}

// Field types
const (
	Numeric = 0
	String  = 1
	Bool    = 2
)

// Operators
const (
	Equal    = 0
	Less     = 1
	Greater  = 2
	Like     = 3
	NotEqual = 4
	Between  = 5
	NotNull  = 6
	Null     = 7
)

// wrap placeholder if field type is string
func getWrappedPlaceholder(fieldType int, arg string) string {
	if fieldType == String {
		return "'" + arg + "'"
	} else {
		return arg
	}
}

func convertFilterDataToSql(a FilterData) string {
	switch a.Operator {
	case NotNull:
		return "t.\"" + a.FieldName + "\"" + " IS NOT NULL"
	case Null:
		return "t.\"" + a.FieldName + "\"" + " IS NULL"
	case Between:

		if a.FieldType != Numeric {
			// Skip wrong - only for numeric type is available
			return ""
		}

		return fmt.Sprintf("t.\"%s\" BETWEEN %s AND %s", a.FieldName,
			getWrappedPlaceholder(a.FieldType, a.Arg0), getWrappedPlaceholder(a.FieldType, a.Arg1))
	case Like:
		if a.FieldType != String {
			// Skip wrong - only for string type is available
			return ""
		}
		return "t.\"" + a.FieldName + "\"" + " LIKE '%" + a.Arg0 + "%'"
	case NotEqual:
		return "t.\"" + a.FieldName + "\"" + " <> " + getWrappedPlaceholder(a.FieldType, a.Arg0)
	case Greater, Less:
		if a.FieldType != Numeric {
			// Skip wrong - only for numeric type is available
			return ""
		}
		operator := " "
		if a.Operator == Less {
			operator = "<"
		} else {
			operator = ">"
		}
		return "t.\"" + a.FieldName + "\"" + " " + operator + " " + getWrappedPlaceholder(a.FieldType, a.Arg0)
	default:
		return "t.\"" + a.FieldName + "\"" + " = " + getWrappedPlaceholder(a.FieldType, a.Arg0)
	}
	return ""
}
