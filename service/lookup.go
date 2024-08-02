package service

import (
	"geo-jot/models"
	"geo-jot/repository"
)

func Find() []models.VehicleWithNearestParcel {
	repoA := repository.NewVehicleRepository()
	results, _ := repoA.GetVehiclesWithNearestParcel()
	return results
}
