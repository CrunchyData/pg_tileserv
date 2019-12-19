package main

import (
	"errors"
	"fmt"

	"net/http"
)

type layerType int

const (
	layerTypeTable    = 1
	layerTypeFunction = 2
)

func (lt layerType) String() string {
	switch lt {
	case layerTypeTable:
		return "table"
	case layerTypeFunction:
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
	GetType() layerType
	GetId() string
	GetDescription() string
	GetName() string
	GetSchema() string
	GetTileRequest(tile Tile, r *http.Request) TileRequest
	WriteLayerJson(w http.ResponseWriter, req *http.Request) error
}

type TileRequest struct {
	Tile Tile
	Sql  string
	Args []interface{}
}

// A global array of Layer where the state is held for performance
// Refreshed when LoadLayerTableList is called
// Key is of the form: schemaname.tablename
var globalLayers map[string]Layer

func GetLayer(lyrId string) (Layer, error) {
	if lyr, ok := globalLayers[lyrId]; ok {
		return lyr, nil
	} else {
		return lyr, errors.New(fmt.Sprintf("Unable to get layer '%s'", lyrId))
	}
}

func LoadLayers() error {
	tableLayers, errTl := GetTableLayers()
	if errTl != nil {
		return errTl
	}
	functionLayers, errFl := GetFunctionLayers()
	if errFl != nil {
		return errFl
	}
	// Empty the global layer map
	globalLayers = make(map[string]Layer)
	for _, lyr := range tableLayers {
		globalLayers[lyr.GetId()] = lyr
	}
	for _, lyr := range functionLayers {
		globalLayers[lyr.GetId()] = lyr
	}

	return nil
}

type LayerJson struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Description string `json:"description"`
	DetailUrl   string `json:"detailurl"`
}

func GetJsonLayers(r *http.Request) map[string]LayerJson {
	json := make(map[string]LayerJson)
	urlBase := serverURLBase(r)
	for k, v := range globalLayers {
		json[k] = LayerJson{
			Type:        v.GetType().String(),
			Id:          v.GetId(),
			Name:        v.GetName(),
			Schema:      v.GetSchema(),
			Description: v.GetDescription(),
			DetailUrl:   fmt.Sprintf("%s/%s.json", urlBase, v.GetId()),
		}
	}
	return json
}
