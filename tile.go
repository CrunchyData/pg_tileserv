package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Tile represents a single tile in a tile
// pyramid, usually referenced in a URL path
// of the form "Zoom/X/Y.Ext"
type Tile struct {
	Zoom int    `json:"zoom"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Ext  string `json:"ext"`
}

// makeTile uses the map populated by the mux.Router
// containing x, y and z keys, and extracts integers
// from them
func makeTile(vars map[string]string) (Tile, error) {
	// Router path restriction ensures
	// these are always numbers
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])
	zoom, _ := strconv.Atoi(vars["z"])
	ext := vars["ext"]
	tile := Tile{Zoom: zoom, X: x, Y: y, Ext: ext}
	// No tile numbers outside the tile grid implied
	// by the zoom?
	if !tile.IsValid() {
		return tile, errors.New(fmt.Sprintf("invalid tile address %s", tile.String()))
	}
	return tile, nil
}

func (tile *Tile) Width() float64 {
	worldTileSize := int(1) << uint(tile.Zoom)
	return worldMercWidth / float64(worldTileSize)
}

// IsValid tests that the tile contains
// only tile addresses that fit within the
// zoom level, and that the zoom level is
// not crazy large
func (tile *Tile) IsValid() bool {
	if tile.Zoom > 32 || tile.Zoom < 0 {
		return false
	}
	worldTileSize := int(1) << tile.Zoom
	if tile.X < 0 || tile.X >= worldTileSize ||
		tile.Y < 0 || tile.Y >= worldTileSize {
		return false
	}
	return true
}

// Bounds calculates the web mercator bounds that
// correspond to this tile
func (tile *Tile) Bounds() Bounds {
	worldMercMax := worldMercWidth / 2
	worldMercMin := -1 * worldMercMax

	// Tile width in EPSG:3857
	tileMercSize := tile.Width()

	// Calculate geographic bounds from tile coordinates
	// XYZ tile coordinates are in "image space" so origin is
	// top-left, not bottom right
	xmin := worldMercMin + (tileMercSize * float64(tile.X))
	xmax := worldMercMin + (tileMercSize * float64(tile.X+1))
	ymin := worldMercMax - (tileMercSize * float64(tile.Y+1))
	ymax := worldMercMax - (tileMercSize * float64(tile.Y))

	return Bounds{xmin, ymin, xmax, ymax}
}

// String returns a path-like representation of the Tile
func (tile *Tile) String() string {
	return fmt.Sprintf("%d/%d/%d.%s", tile.Zoom, tile.X, tile.Y, tile.Ext)
}
