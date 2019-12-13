package main

type layerType int

const (
	layerTypeTable    = 1
	layerTypeFunction = 2
)

// type Layer interface {
// 	GetType() layerType
// 	GetId() string
// 	GetDescription() string
// 	GetName() string
// 	GetSchema() string
// 	GetUrl(urlBase string) string()
// }

/*

R http.Request

GetLayers(R) => []LayerJson
GetLayerDetail(R, Id) => LayerDetailJson
GetTile(R, Id, ZXY) 
  
  GetTileRequest(Tile, R, Layer, )





*/