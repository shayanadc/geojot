package test

import (
	"context"
	"geo-jot/models"
	"geo-jot/repository"
	"testing"
)

func TestGetAll(t *testing.T) {
	_, _ = SetupTestDatabase(t)

	vehicles := []models.Vehicle{
		{Number: "ABC123", Coordinates: []float64{40.7128, -74.0060}},
		{Number: "XYZ789", Coordinates: []float64{34.0522, -118.2437}},
	}

	var documents []interface{}
	for _, v := range vehicles {
		documents = append(documents, v)
	}

	repo := repository.NewVehicleRepository()

	_, err := repo.Collection.InsertMany(context.Background(), documents)

	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	result, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll() returned an error: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Expected 2 vehicles in the result, got %d", len(result))
	}

}
