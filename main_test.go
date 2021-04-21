package main

import (
	"context"
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

	// viper.Set("DbConnection", os.Getenv("TEST_DATABASE_URL"))
	viper.Set("DbConnection", os.Getenv("dbname=ts"))
	db, err := dbConnect()
	if err != nil {
		os.Exit(1)
	}
	sql := "CREATE EXTENSION IF NOT EXISTS postgis"
	_, err = db.Exec(context.Background(), sql)
	if err != nil {

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
	}

}
