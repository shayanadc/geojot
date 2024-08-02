package test

import (
	"fmt"
	"geo-jot/config"
	"geo-jot/container"
	"geo-jot/db"
	"testing"
)

func SetupTestDatabase(t *testing.T) (*db.Client, func()) {
	config.LoadEnv("../.env.test")
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
