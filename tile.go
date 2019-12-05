package main

import (
	// log "github.com/sirupsen/logrus"
)

type Tile struct {
	Zoom int    `json:"zoom"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Ext  string `json:"ext"`
}

// A global constant, the width of the Web Mercator plane
const worldMercWidth float64 = 40075016.6855784

func (tile *Tile) Width() float64 {
	worldTileSize := int(1) << tile.Zoom
	return worldMercWidth / float64(worldTileSize)
}

func (tile *Tile) IsValid() bool {
	worldTileSize := int(1) << tile.Zoom
	if tile.X < 0 || tile.X >= worldTileSize ||
		tile.Y < 0 || tile.Y >= worldTileSize {
		return false
	}
	return true
}

func (tile *Tile) Bounds() Bounds {
	worldMercMax := worldMercWidth / 2
	worldMercMin := -1 * worldMercMax

	// Tile width in EPSG:3857
	tileMercSize := tile.Width()

	// Calculate geographic bounds from tile coordinates
	// XYZ tile coordinates are in "image space" so origin is
	// top-left, not bottom right
	xmin := worldMercMin + tileMercSize*float64(tile.X)
	xmax := worldMercMin + tileMercSize*float64(tile.X+1)
	ymin := worldMercMax - tileMercSize*float64(tile.Y+1)
	ymax := worldMercMax - tileMercSize*float64(tile.Y)

	return Bounds{xmin, ymin, xmax, ymax}
}
