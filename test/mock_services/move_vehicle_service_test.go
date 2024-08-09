package mock_services_test

import (
	"geo-jot/models"
	"geo-jot/service"
	"testing"
)

type MockVehicle struct {
	position int
}

func (v *MockVehicle) Move(distance int) {
	v.position += distance
}

func TestMoveVehicles(t *testing.T) {

	vehicles := []models.Vehicle{
		{Number: "ABC123", Coordinates: []float64{40.7128, -74.0060}},
		{Number: "ABC123", Coordinates: []float64{40.6119, -74.10005}},
	}
	mockRepo := &MockVehicleRepository{vehicles: vehicles}

	newVehicles := service.MoveVehicles(mockRepo)

	if len(vehicles) != 2 {
		t.Errorf("Expected 2 vehicles, got %d", len(vehicles))
	}

	if newVehicles[0].Coordinates[0] == 40.7128 {
		t.Errorf("Expected first vehicle to have coordinates [40.7128, -74.0060], got %v", vehicles[0].Coordinates)
	}

	if newVehicles[1].Coordinates[1] == -74.10005 {
		t.Errorf("Expected first vehicle to have coordinates [40.7128, -74.0060], got %v", vehicles[0].Coordinates)
	}
}
