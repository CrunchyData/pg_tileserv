package main

import (
	"testing"
)

func TestParseArgDefault(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"foo", "foo"},
		{"123", "123"},
		{"'-123'::integer", "-123"},
		{"'-123.1'::numeric", "-123.1"},
		{"'foo'::text", "foo"},
		{"'foo::bar'::text", "foo::bar"},
	}

	for _, test := range tests {
		if output := parseArgDefault(test.input); output != test.output {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.output, output)
		}
	}

}
