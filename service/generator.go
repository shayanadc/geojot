package service

import (
	"geo-jot/models"
	"geo-jot/repository"
	"strconv"

	"golang.org/x/exp/rand"
)

func GenerateRandomVehicle() models.Vehicle {
	vehicleNumber := "V" + strconv.Itoa(rand.Intn(10000))

	latitude := rand.Float64()*180 - 90   // Latitude ranges from -90 to 90
	longitude := rand.Float64()*360 - 180 // Longitude ranges from -180 to 180

	return models.Vehicle{
		Number:      vehicleNumber,
		Coordinates: []float64{latitude, longitude},
	}
}

func GenerateRandomVehicles(count int) []models.Vehicle {
	vehicles := make([]models.Vehicle, count)
	for i := 0; i < count; i++ {
		vehicles[i] = GenerateRandomVehicle()
	}
	return vehicles
}

func StoreMany() {
	repo := repository.NewVehicleRepository()
	vehicles := GenerateRandomVehicles(50)

	_ = repo.InsertMany(vehicles)
}
