package service

import (
	"geo-jot/models"
	"geo-jot/repository"
)

func MoveVehicles(repo repository.VehicleRepository) []models.Vehicle {

	vehicles, _ := repo.GetLatest()
	for _, vehicle := range vehicles {
		vehicle.Move(10)
	}
	return vehicles
}
