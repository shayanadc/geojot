package mock_services_test

import (
	"geo-jot/models"
)

type MockVehicleRepository struct {
	vehicles []models.Vehicle
}

func (m *MockVehicleRepository) GetLatest() ([]models.Vehicle, error) {
	return m.vehicles, nil
}
