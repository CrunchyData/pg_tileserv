package main

import (
	"encoding/json"
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

func getLayer(lyrID string) (Layer, error) {
	lyr, ok := globalLayers[lyrID]
	if ok {
		return lyr, nil
	}
	return lyr, fmt.Errorf("Unable to get layer '%s'", lyrID)
}

func loadLayers() error {
	_, errBnd := getServerBounds()
	if errBnd != nil {
		return errBnd
	}
	tableLayers, errTl := getTableLayers()
	if errTl != nil {
		return errTl
	}
	functionLayers, errFl := getFunctionLayers()
	if errFl != nil {
		return errFl
	}
	// Empty the global layer map
	globalLayersMutex.Lock()
	globalLayers = make(map[string]Layer)
	for _, lyr := range tableLayers {
		globalLayers[lyr.GetID()] = lyr
	}
	for _, lyr := range functionLayers {
		globalLayers[lyr.GetID()] = lyr
	}
	globalLayersMutex.Unlock()

	return nil
}

type layerJSON struct {
	Type        string                 `json:"type"`
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Schema      string                 `json:"schema"`
	Description string                 `json:"description,omitempty"`
	DetailURL   string                 `json:"detailurl"`
	ExtraFields map[string]interface{} `json:"-"`
}

func (lj layerJSON) MarshalJSON() ([]byte, error) {
	type Alias layerJSON
	base := Alias(lj)
	b, _ := json.Marshal(base)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range lj.ExtraFields {
		m[k] = v
	}
	return json.Marshal(m)
}

// Main function that reads JSON from Description and create the additional fields
func getJSONLayers(r *http.Request) map[string]layerJSON {
	jsonLayers := make(map[string]layerJSON)
	urlBase := serverURLBase(r)
	globalLayersMutex.Lock()
	defer globalLayersMutex.Unlock()
	for k, v := range globalLayers {
		lj := layerJSON{
			Type:      v.GetType().String(),
			ID:        v.GetID(),
			Name:      v.GetName(),
			Schema:    v.GetSchema(),
			DetailURL: fmt.Sprintf("%s/%s.json", urlBase, url.PathEscape(v.GetID())),
		}
		var descFields map[string]interface{}
		err := json.Unmarshal([]byte(v.GetDescription()), &descFields)
		if err != nil {
			lj.Description = v.GetDescription()
		} else {
			lj.ExtraFields = descFields
		}
		jsonLayers[k] = lj
	}
	return jsonLayers
}
