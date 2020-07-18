package models

// Coordinates GeoJSON
type Coordinates []float64

// Location GeoJSON
type Location struct {
	Type        *string      `json:"type,omitempty" bson:"type"`
	Coordinates *Coordinates `json:"coordinates,omitempty" bson:"coordinates"`
}

// RidesSequence - Sequence struct
type RidesSequence struct {
	Sequence int `bson:"sequence"`
}
