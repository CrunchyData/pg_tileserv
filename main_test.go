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
	"github.com/stretchr/testify/suite"
)

type dbSuite struct {
	suite.Suite
	setup bool
}

var db *pgxpool.Pool

func (suite *dbSuite) SetupSuite() {
	viper.Set("DbConnection", os.Getenv("TEST_DATABASE_URL"))
	db, err := DbConnect()
	if err != nil {
		suite.setup = false
		return
	}
	sql := "CREATE EXTENSION IF NOT EXISTS postgis"
	_, err = db.Exec(context.Background(), sql)
	if err != nil {
		suite.T().Skip("DB integration test suite setup failed, skipping")
		suite.setup = false
		return
	}
	suite.setup = true
}

func (suite *dbSuite) TearDownSuite() {
}

func (suite *dbSuite) SetupTest() {
}

func (suite *dbSuite) TearDownTest() {
}

func (suite *dbSuite) TestDBNoTables() {
	if !suite.setup {
		suite.T().Skip("DB integration test suite setup failed, skipping")
	}
	r := TileRouter()
	request, _ := http.NewRequest("GET", "/index.json", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(suite.T(), 200, response.Code, "OK response is expected")

	json_result := response.Body.String()
	json_result = strings.TrimSpace(json_result)
	json_expect := "{}"
	assert.Equal(suite.T(), json_expect, json_result, "empty json response is expected")
}

func TestDatabaseSuite(t *testing.T) {
	tests := new(dbSuite)
	suite.Run(t, tests)
}
