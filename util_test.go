package main

import (
	"testing"
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
			[]string{"http://example.com/", "/tiles"},
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
