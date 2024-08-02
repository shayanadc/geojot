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

func GetAll() []models.Vehicle {
	repoA := repository.NewVehicleRepository()
	results, _ := repoA.GetAll()
	return results
}

func GetLatest() []models.Vehicle {
	repoA := repository.NewVehicleRepository()
	results, _ := repoA.GetLatest()
	return results
}
