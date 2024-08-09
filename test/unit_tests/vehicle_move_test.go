package models

import (
	"geo-jot/models"
	"testing"
)

func TestVehicle_Move(t *testing.T) {
	lat := 40.7128
	lon := -74.006
	vehicle := models.Vehicle{
		Number:      "ABC123",
		Coordinates: []float64{lat, lon},
	}

	vehicle.Move(10)

	if vehicle.Coordinates[0] == lat {
		t.Errorf("Vehicle coordinates after move are incorrect. Got %v", vehicle.Coordinates[0])
	}

	if vehicle.Coordinates[1] == lon {
		t.Errorf("Vehicle coordinates after move are incorrect. Got %v", vehicle.Coordinates[1])
	}
}
