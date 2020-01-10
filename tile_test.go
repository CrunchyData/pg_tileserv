package main

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

func Test_makeTile(t *testing.T) {
	// func makeTile(vars map[string]string) (Tile, error) {

	var tests = []struct {
		input   map[string]string
		outtile Tile
		outerr  error
	}{
		{
			map[string]string{"x": "0", "y": "0", "z": "0", "ext": "pbf"},
			Tile{X: 0, Y: 0, Zoom: 0, Ext: "pbf"},
			nil,
		},
	}

	for _, test := range tests {
		if output, err := makeTile(test.input); test.outtile != output || test.outerr != err {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}, {}", test.input, test.outtile, output, err)
		}
	}

}

func Test_tileIsValid(t *testing.T) {
	// func makeTile(vars map[string]string) (Tile, error) {

	var tests = []struct {
		input  Tile
		output bool
	}{
		{Tile{X: 0, Y: 0, Zoom: 0, Ext: "pbf"}, true},
		{Tile{X: 0, Y: 1, Zoom: 0, Ext: "pbf"}, false},
		{Tile{X: 1, Y: 1, Zoom: 1, Ext: "pbf"}, true},
		{Tile{X: -1, Y: 0, Zoom: 0, Ext: "pbf"}, false},
		{Tile{X: 0, Y: 0, Zoom: -1, Ext: "pbf"}, false},
	}

	for _, test := range tests {
		if output := test.input.IsValid(); output != test.output {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.output, output)
		}
	}

}
