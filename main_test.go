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

	viper.Set("DbConnection", os.Getenv("TEST_DATABASE_URL"))
	db, err := DbConnect()
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
	r := TileRouter()
	request, _ := http.NewRequest("GET", "/index.json", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

	json_result := strings.TrimSpace(response.Body.String())
	json_expect := "{}"
	assert.Equal(t, json_expect, json_result, "empty json response is expected")
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
		r := TileRouter()
		request, _ := http.NewRequest("GET", "/test/index.json", nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)
		assert.Equal(t, 200, response.Code, "OK response is expected")
	}

}
