package models

import "time"

// Fare - ride fare
type Fare struct {
	ChargesPerKM  float64 `json:"charges_per_km" bson:"charges_per_km"`
	ChargesPerMin float64 `json:"charges_per_min" bson:"charges_per_min"`
}

// Ride - Ride model
type Ride struct {
	ID           *int      `json:"id,omitempty" bson:"_id,omitempty" valid:"required"`
	Rider        *string   `json:"rider,omitempty" bson:"rider,omitempty" valid:"required"`
	Provider     *string   `json:"provider,omitempty" bson:"provider,omitempty" valid:"required"`
	StartTime    time.Time `json:"start" bson:"start"`
	ModifiedTime time.Time `json:"modified" bson:"modified"`
	Status       int       `json:"status,omitempty" bson:"status" valid:"required"`
	Pickup       *Location `json:"pickup,omitempty" bson:"pickup" valid:"required"`
	Drop         *Location `json:"drop,omitempty" bson:"drop" valid:"required"`
	Fare         Fare      `json:"fare"  bson:"fare"`
}
