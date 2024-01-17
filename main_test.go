package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var db *pgxpool.Pool
var dbsetup = false

func TestMain(m *testing.M) {

	viper.Set("DbConnection", os.Getenv("TEST_DATABASE_URL"))
	// viper.Set("DbConnection", os.Getenv("dbname=ts"))
	db, err := dbConnect()
	if err != nil {
		os.Exit(1)
	}
	sql := "CREATE EXTENSION IF NOT EXISTS postgis"
	_, err = db.Exec(context.Background(), sql)
	if err != nil {
		fmt.Printf("Error creating extension: %s", err)
		os.Exit(1)
	}

	dbsetup = true
	os.Exit(m.Run())
}

func TestDBNoTables(t *testing.T) {
	if !dbsetup {
		t.Skip("DB integration test suite setup failed, skipping")
	}
	r := tileRouter()
	request, _ := http.NewRequest("GET", "/index.json", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

	jsonResult := strings.TrimSpace(response.Body.String())
	jsonExpect := "{}"
	assert.Equal(t, jsonExpect, jsonResult, "empty json response is expected")
}

// TestBasePath sets an alternate base path to check that handlers are
// mounted at the specified path
func TestBasePath(t *testing.T) {

	if !dbsetup {
		t.Skip("DB integration test suite setup failed, skipping")
	}

	// paths to check
	paths := []string{"/test", "/test/"}

	for _, path := range paths {
		viper.Set("BasePath", path)
		r := tileRouter()
		request, _ := http.NewRequest("GET", "/test/index.json", nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")

		request, _ = http.NewRequest("GET", "/test/health", nil)
		response = httptest.NewRecorder()
		r.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	}

	// cleanup
	viper.Set("BasePath", "/")
}

// Test that the preview endpoints are hidden or shown according to the config
func TestShowPreview(t *testing.T) {
	viper.Set("ShowPreview", true)
	r := tileRouter()
	request, _ := http.NewRequest("GET", "/index.json", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	request, _ = http.NewRequest("GET", "/index.html", nil)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

// the current default behavior is to show the preview
func TestShowPreviewDefault(t *testing.T) {
	r := tileRouter()
	request, _ := http.NewRequest("GET", "/index.json", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	request, _ = http.NewRequest("GET", "/index.html", nil)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestHidePreview(t *testing.T) {
	viper.Set("ShowPreview", false)
	r := tileRouter()
	request, _ := http.NewRequest("GET", "/index.json", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "Not Found response is expected")
	request, _ = http.NewRequest("GET", "/index.html", nil)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "Not Found response is expected")

	// cleanup
	viper.Set("ShowPreview", true)
}

// Test that the health endpoint gives a 200 if the server is running
func TestHealth(t *testing.T) {
	r := tileRouter()
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "200 OK", string(response.Result().Status), "Response status should say ok")
}

func TestHealthCustomUrl(t *testing.T) {
	viper.Set("HealthEndpoint", "/testHealthABC")
	r := tileRouter()
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code, "Not Found response is expected")
	request, _ = http.NewRequest("GET", "/testHealthABC", nil)
	response = httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "200 OK", string(response.Result().Status), "Response status should say ok")

	// cleanup
	viper.Set("HealthEndpoint", "/health")
}
