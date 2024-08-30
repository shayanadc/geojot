package service

import (
	"geo-jot/models"
	"geo-jot/repository"
)

type LookupService struct {
	vehicleRepo repository.VehicleRepository
}

func NewLookupService(vehicleRepo repository.VehicleRepository) *LookupService {
	return &LookupService{
		vehicleRepo: vehicleRepo,
	}
}

func (service LookupService) Find() []models.VehicleWithNearestParcel {
	results, _ := service.vehicleRepo.GetVehiclesWithNearestParcel()
	return results
}

func (service LookupService) GetAll() []models.Vehicle {
	results, _ := service.vehicleRepo.GetAll()
	return results
}

func (service LookupService) GetLatest() []models.Vehicle {
	results, _ := service.vehicleRepo.GetLatest()
	return results
}
