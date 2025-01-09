package main

import (
	"fmt"

	"net/http"
	"net/url"
)

// LayerType is the table/function type of a layer
type LayerType int

const (
	// LayerTypeTable is a table layer
	LayerTypeTable = 1
	// LayerTypeFunction is a function layer
	LayerTypeFunction = 2
)

func (lt LayerType) String() string {
	switch lt {
	case LayerTypeTable:
		return "table"
	case LayerTypeFunction:
		return "function"
	default:
		return "unknown"
	}
}

// A Layer is a LayerTable or a LayerFunction
// in either case it should be able to generate
// a TileRequest containing SQL to produce tiles
// given an input tile
type Layer interface {
	GetType() LayerType
	GetID() string
	GetDescription() string
	GetName() string
	GetSchema() string
	GetTileRequest(tile Tile, r *http.Request) TileRequest
	WriteLayerJSON(w http.ResponseWriter, req *http.Request) error
}

// A TileRequest specifies what to fetch from the database for a single tile
type TileRequest struct {
	LayerID string
	Tile    Tile
	SQL     string
	Args    []interface{}
}

func getLayer(lyrID string, databaseRole string) (Layer, error) {

	// check if layers have already been listed for this databaseRole
	_, ok := globalLayers[databaseRole]
	if !ok {
		loadLayers(databaseRole)
	}

	lyr, ok := globalLayers[databaseRole][lyrID]
	if ok {
		return lyr, nil
	}
	return lyr, fmt.Errorf("Unable to get layer '%s'", lyrID)
}

func loadLayers(databaseRole string) error {
	_, errBnd := getServerBounds()
	if errBnd != nil {
		return errBnd
	}
	tableLayers, errTl := getTableLayers(databaseRole)
	if errTl != nil {
		return errTl
	}
	functionLayers, errFl := getFunctionLayers(databaseRole)
	if errFl != nil {
		return errFl
	}
	// Empty the global layer map for this databaseRole only
	globalLayersMutex.Lock()
	globalLayers[databaseRole] = make(map[string]Layer)
	for _, lyr := range tableLayers {
		globalLayers[databaseRole][lyr.GetID()] = lyr
	}
	for _, lyr := range functionLayers {
		globalLayers[databaseRole][lyr.GetID()] = lyr
	}
	globalLayersMutex.Unlock()

	return nil
}

type layerJSON struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Description string `json:"description"`
	DetailURL   string `json:"detailurl"`
}

func getJSONLayers(r *http.Request, databaseRole string) map[string]layerJSON {
	// check if layers have already been listed for this databaseRole
	_, ok := globalLayers[databaseRole]
	if !ok {
		loadLayers(databaseRole)
	}

	json := make(map[string]layerJSON)
	urlBase := serverURLBase(r)
	globalLayersMutex.Lock()
	for k, v := range globalLayers[databaseRole] {
		json[k] = layerJSON{
			Type:        v.GetType().String(),
			ID:          v.GetID(),
			Name:        v.GetName(),
			Schema:      v.GetSchema(),
			Description: v.GetDescription(),
			DetailURL:   fmt.Sprintf("%s/%s.json", urlBase, url.PathEscape(v.GetID())),
		}
	}
	globalLayersMutex.Unlock()
	return json
}
