package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_formatBaseURL(t *testing.T) {
	var tests = []struct {
		input  []string
		output string
	}{
		{
			[]string{"http://example.com", "/tiles"},
			"http://example.com/tiles",
		},
		{
			[]string{"http://example.com", "tiles"},
			"http://example.com/tiles",
		},
		{
			[]string{"http://example.com/", "/tiles/"},
			"http://example.com/tiles",
		},
		{
			[]string{"https://example.com/", "/"},
			"https://example.com",
		},
	}

	for _, test := range tests {
		if output := formatBaseURL(test.input[0], test.input[1]); output != test.output {
			t.Errorf("Test failed: input: %v, expected: %s, recieved: %s", test.input, test.output, output)
		}
	}
}

func TestMetrics(t *testing.T) {

	if !dbsetup {
		t.Skip("DB integration test suite setup failed, skipping")
	}

	viper.Set("EnableMetrics", true)

	r := tileRouter()
	request, _ := http.NewRequest("GET", "/metrics", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
