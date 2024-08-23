package service

import (
	"fmt"
	"geo-jot/models"
	"geo-jot/repository"
	"strconv"
	"time"

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

func StoreConcurrently() {
	repo := repository.NewVehicleRepository()
	vehicles := GenerateRandomVehicles(80000)

	vehiclesA := vehicles[:20000]
	vehiclesB := vehicles[20000:40000]
	vehiclesC := vehicles[40000:60000]
	vehiclesD := vehicles[60000:80000]

	now := time.Now()

	go func(vehicles []models.Vehicle) {
		_ = repo.InsertMany(vehicles)
	}(vehiclesA)

	go func(vehicles []models.Vehicle) {
		_ = repo.InsertMany(vehicles)
	}(vehiclesB)

	go func(vehicles []models.Vehicle) {
		_ = repo.InsertMany(vehicles)
	}(vehiclesC)

	go func(vehicles []models.Vehicle) {
		_ = repo.InsertMany(vehicles)
	}(vehiclesD)

	fmt.Println("Time taken:", time.Since(now))

}
