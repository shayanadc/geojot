package models

import (
	"math"
)

type Vehicle struct {
	Number      string    `bson:"number"`
	Coordinates []float64 `bson:"coordinates"`
}

func (v *Vehicle) Move(distanceMeters float64) {
	const earthRadius = 6371000.0 // Earth's radius in meters
	const degToRad = math.Pi / 180.0
	const radToDeg = 180.0 / math.Pi
	// Current latitude and longitude
	lat := v.Coordinates[0]
	lon := v.Coordinates[1]

	// Move latitude (North-South)
	newLat := lat + (distanceMeters/earthRadius)*radToDeg

	// Move longitude (East-West)
	// Adjust longitude based on the latitude
	newLon := lon + (distanceMeters/(earthRadius*math.Cos(lat*degToRad)))*radToDeg

	// Update the coordinates with the new position
	v.Coordinates[0] = newLat
	v.Coordinates[1] = newLon
}
