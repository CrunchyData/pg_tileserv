package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Tile represents a single tile in a tile
// pyramid, usually referenced in a URL path
// of the form "Zoom/X/Y.Ext"
type Tile struct {
	Zoom   int    `json:"zoom"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Ext    string `json:"ext"`
	Bounds Bounds `json:"bounds"`
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
		invalidTileError := tileAppError{
			HTTPCode: 400,
			SrcErr:   errors.New(fmt.Sprintf("invalid tile address %s", tile.String())),
		}
		return tile, invalidTileError
	}
	e := tile.CalculateBounds()
	return tile, e
}

func (tile *Tile) width() float64 {
	return math.Abs(tile.Bounds.Xmax - tile.Bounds.Xmin)
}

// IsValid tests that the tile contains
// only tile addresses that fit within the
// zoom level, and that the zoom level is
// not crazy large
func (tile *Tile) IsValid() bool {
	if tile.Zoom > 32 || tile.Zoom < 0 {
		return false
	}
	worldTileSize := int(1) << uint(tile.Zoom)
	if tile.X < 0 || tile.X >= worldTileSize ||
		tile.Y < 0 || tile.Y >= worldTileSize {
		return false
	}
	return true
}

// CalculateBounds calculates the cartesian bounds that
// correspond to this tile
func (tile *Tile) CalculateBounds() (e error) {
	serverBounds, e := getServerBounds()
	if e != nil {
		return e
	}

	worldWidthInTiles := float64(int(1) << uint(tile.Zoom))
	tileWidth := math.Abs(serverBounds.Xmax-serverBounds.Xmin) / worldWidthInTiles

	// Calculate geographic bounds from tile coordinates
	// XYZ tile coordinates are in "image space" so origin is
	// top-left, not bottom right
	xmin := serverBounds.Xmin + (tileWidth * float64(tile.X))
	xmax := serverBounds.Xmin + (tileWidth * float64(tile.X+1))
	ymin := serverBounds.Ymax - (tileWidth * float64(tile.Y+1))
	ymax := serverBounds.Ymax - (tileWidth * float64(tile.Y))
	tile.Bounds = Bounds{serverBounds.SRID, xmin, ymin, xmax, ymax}

	return nil
}

// String returns a path-like representation of the Tile
func (tile *Tile) String() string {
	return fmt.Sprintf("%d/%d/%d.%s", tile.Zoom, tile.X, tile.Y, tile.Ext)
}
