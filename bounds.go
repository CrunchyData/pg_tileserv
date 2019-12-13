package main

import (
	"fmt"
	"math"
)

type Bounds struct {
	Xmin float64 `json:"Xmin"`
	Ymin float64 `json:"Ymin"`
	Xmax float64 `json:"Xmax"`
	Ymax float64 `json:"Ymax"`
}

func (b *Bounds) String() string {
	return fmt.Sprintf("{Xmin:%g, Ymin:%g, Xmax:%g, Ymax:%g}",
		b.Xmin, b.Ymin, b.Xmax, b.Ymax)
}

func (b *Bounds) Sql() string {
	return fmt.Sprintf("ST_MakeEnvelope(%g, %g, %g, %g, 3857)",
		b.Xmin, b.Ymin,
		b.Xmax, b.Ymax)
}

func (b *Bounds) Expand(size float64) {
	worldMin := -0.5 * worldMercWidth
	worldMax := 0.5 * worldMercWidth
	b.Xmin = math.Max(b.Xmin-size, worldMin)
	b.Ymin = math.Max(b.Ymin-size, worldMin)
	b.Xmax = math.Min(b.Xmax+size, worldMax)
	b.Ymax = math.Min(b.Ymax+size, worldMax)
	return
}
