package service

import (
	"geo-jot/models"
	"geo-jot/repository"
)

type MoverService struct {
	vehicleRepo repository.VehicleRepository
}

func NewMoverService(vehicleRepo repository.VehicleRepository) *MoverService {
	return &MoverService{
		vehicleRepo: vehicleRepo,
	}
}

func (service MoverService) MoveVehicles() []models.Vehicle {
	vehicles, _ := service.vehicleRepo.GetLatest()

	for _, vehicle := range vehicles {
		vehicle.Move(10)
	}
	return vehicles
}

func (service MoverService) InsertVehiclesMove() {

	vehicles := service.MoveVehicles()

	_ = service.vehicleRepo.InsertMany(vehicles)
}
