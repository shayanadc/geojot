package test

import (
	"fmt"
	"geo-jot/config"
	"geo-jot/container"
	"geo-jot/db"
	"geo-jot/models"
	"testing"
)

type MockVehicleRepository struct {
	Vehicles []models.Vehicle
}

func (m *MockVehicleRepository) GetLatest() ([]models.Vehicle, error) {
	return m.Vehicles, nil
}
func (m *MockVehicleRepository) InsertMany(vehicles []models.Vehicle) error {
	return nil
}

func SetupTestDatabase(t *testing.T) (*db.Client, func()) {
	_ = config.LoadEnv("../../.env.test")
	conn := db.NewDatabaseConnection()
	fmt.Printf(conn.Name)
	dbClient := db.NewClient(conn)
	container.GetContainer().SetDBClient(dbClient)

	cleanup := func() {
		if err := dbClient.DropDatabase(); err != nil {
			t.Fatalf("Failed to drop test database: %v", err)
		}
	}

	t.Cleanup(cleanup)

	return dbClient, cleanup
}
