package service

import (
	"geo-jot/models"
	"math"
	"math/rand"
)

const earthRadius = 6371000

func MoveVehicle(vehicle *models.Vehicle, distanceInMeters float64, rng *rand.Rand) {
	lat := vehicle.Coordinates[0]
	lon := vehicle.Coordinates[1]

	angle := rng.Float64() * 2 * math.Pi

	distanceInRadians := distanceInMeters / earthRadius

	newLat := lat + (distanceInRadians * math.Cos(angle) * (180 / math.Pi))

	newLon := lon + (distanceInRadians*math.Sin(angle)*(180/math.Pi))/math.Cos(lat*math.Pi/180)

	vehicle.Coordinates[0] = newLat
	vehicle.Coordinates[1] = newLon
}
