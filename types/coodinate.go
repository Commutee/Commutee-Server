package types

type FeatureCollection struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
