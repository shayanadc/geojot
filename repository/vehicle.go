package repository

import (
	"context"
	"fmt"
	"geo-jot/container"
	"geo-jot/db"
	"geo-jot/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VehicleRepository interface {
	GetLatest() ([]models.Vehicle, error)
}

const vehicleCollection string = "vehicles"

type vehicleRepository struct {
	DB         *db.Client
	Collection *mongo.Collection
}

func NewVehicleRepository() *vehicleRepository {
	db := container.GetContainer().GetDBClient()

	return &vehicleRepository{DB: db, Collection: db.GetCollection(vehicleCollection)}
}
func (repo *vehicleRepository) GetLatest() ([]models.Vehicle, error) {

	pipeline := mongo.Pipeline{
		{{"$sort", bson.D{{"number", -1}}}}, // Sort by number and then by timestamp in descending order
		{{"$group", bson.D{
			{"_id", "$number"},
			{"latestDocument", bson.D{{"$last", "$$ROOT"}}}, // Use $first to get the latest document after sorting
		}}},
		{{"$replaceRoot", bson.D{{"newRoot", "$latestDocument"}}}},
	}

	cursor, err := repo.Collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var vehicles []models.Vehicle
	if err = cursor.All(context.Background(), &vehicles); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *vehicleRepository) GetAll() ([]models.Vehicle, error) {
	ctx := context.Background()

	cursor, err := r.Collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to execute find operation: %w", err)
	}
	defer cursor.Close(ctx)

	var vehicles []models.Vehicle
	if err = cursor.All(ctx, &vehicles); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	return vehicles, nil
}
func (r *vehicleRepository) GetVehiclesWithNearestParcel() ([]models.VehicleWithNearestParcel, error) {

	cursor, err := r.Collection.Aggregate(r.DB.Context, bson.A{
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "parcels"},
					{"let", bson.D{{"vehicleLoc", "$loc"}}},
					{"pipeline",
						bson.A{
							bson.D{
								{"$geoNear",
									bson.D{
										{"near", "$$vehicleLoc"},
										{"distanceField", "dist.calculated"},
										{"spherical", true},
									},
								},
							},
							bson.D{{"$limit", 1}},
						},
					},
					{"as", "nearestParcel"},
				},
			},
		},
		bson.D{{"$unwind", "$nearestParcel"}},
		bson.D{
			{"$project",
				bson.D{
					{"number", 1},
					{"vehicleLoc", "$loc"},
					{"nearestParcel",
						bson.D{
							{"_id", "$nearestParcel._id"},
							{"loc", "$nearestParcel.loc"},
							{"distance", "$nearestParcel.dist.calculated"},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(r.DB.Context)

	var results []models.VehicleWithNearestParcel
	if err = cursor.All(r.DB.Context, &results); err != nil {
		log.Fatal(err)
	}
	return results, nil
}
