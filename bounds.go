package main

import (
	"fmt"
	"math"
)

// Bounds represents a box in Web Mercator space
type Bounds struct {
	Xmin float64 `json:"xmin"`
	Ymin float64 `json:"ymin"`
	Xmax float64 `json:"xmax"`
	Ymax float64 `json:"ymax"`
}

func (b *Bounds) String() string {
	return fmt.Sprintf("{Xmin:%g, Ymin:%g, Xmax:%g, Ymax:%g}",
		b.Xmin, b.Ymin, b.Xmax, b.Ymax)
}

// SQL returns the SQL fragment to create this bounds in the database
func (b *Bounds) SQL() string {
	return fmt.Sprintf("ST_MakeEnvelope(%g, %g, %g, %g, 3857)",
		b.Xmin, b.Ymin,
		b.Xmax, b.Ymax)
}

// Expand increases the size of this bounds in all directions, respecting
// the limits of the Web Mercator plane
func (b *Bounds) Expand(size float64) {
	worldMin := -0.5 * worldMercWidth
	worldMax := 0.5 * worldMercWidth
	b.Xmin = math.Max(b.Xmin-size, worldMin)
	b.Ymin = math.Max(b.Ymin-size, worldMin)
	b.Xmax = math.Min(b.Xmax+size, worldMax)
	b.Ymax = math.Min(b.Ymax+size, worldMax)
	return
}

func fromMercator(x float64, y float64) (lng float64, lat float64) {
	mercSize := worldMercWidth / 2.0
	lng = x * 180.0 / mercSize
	lat = 180.0 / math.Pi * (2.0*math.Atan(math.Exp((y/mercSize)*math.Pi)) - math.Pi/2.0)
	return lng, lat
}

// Json returns the bounds in array for form consumption
// by Json formats that like it that way
func (b *Bounds) Json() []float64 {
	s := make([]float64, 4)
	s[0], s[1] = fromMercator(b.Xmin, b.Ymin)
	s[2], s[3] = fromMercator(b.Xmax, b.Ymax)
	return s
}

// Center returns the center of the bounds in array format
// for consumption by Json formats that like it that way
func (b *Bounds) Center() []float64 {
	xc := (b.Xmin + b.Xmax) / 2.0
	yc := (b.Ymin + b.Ymax) / 2.0
	s := make([]float64, 2)
	s[0], s[1] = fromMercator(xc, yc)
	return s
}
