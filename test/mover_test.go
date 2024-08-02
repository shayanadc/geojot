package test

import (
	"geo-jot/models"
	"geo-jot/service"
	"math/rand"
	"testing"
)

const earthRadius = 6371000

func TestMoveVehicle(t *testing.T) {
	fixedSeed := rand.NewSource(42)
	rng := rand.New(fixedSeed)

	vehicle := &models.Vehicle{
		Number:      "TEST123",
		Coordinates: []float64{40.7128, -74.0060},
	}

	service.MoveVehicle(vehicle, 10, rng)
	if vehicle.Coordinates[0] == 40.7128 {
		t.Errorf("Expected latitude %f, but got %f", vehicle.Coordinates[0], 40.7128)

	}

	service.MoveVehicle(vehicle, 10, rng)
	if vehicle.Coordinates[1] == -74.0060 {
		t.Errorf("Expected latitude %f, but got %f", vehicle.Coordinates[1], -74.0060)
	}
}
