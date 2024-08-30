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
	InsertMany([]models.Vehicle) error
	GetVehiclesWithNearestParcel() ([]models.VehicleWithNearestParcel, error)
	GetAll() ([]models.Vehicle, error)
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

func (repo *vehicleRepository) InsertMany(vehicles []models.Vehicle) error {

	docs := make([]interface{}, len(vehicles))
	for i, v := range vehicles {
		docs[i] = v
	}
	_, err := repo.Collection.InsertMany(context.Background(), docs)

	if err != nil {
		log.Printf("Failed to insert vehicles: %v", err)
		return err
	}

	return nil
}

func (repo *vehicleRepository) GetLatest() ([]models.Vehicle, error) {

	pipeline := mongo.Pipeline{
		{{Key: "$sort", Value: bson.D{{Key: "number", Value: -1}}}}, // Sort by number in descending order
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$number"},
			{Key: "latestDocument", Value: bson.D{{Key: "$last", Value: "$$ROOT"}}}, // Use $last to get the latest document after sorting
		}}},
		{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$latestDocument"}}}},
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

	pipeline := mongo.Pipeline{
		{
			{
				Key: "$lookup", Value: bson.D{
					{Key: "from", Value: "parcels"},
					{Key: "let", Value: bson.D{{Key: "vehicleLoc", Value: "$loc"}}},
					{Key: "pipeline", Value: mongo.Pipeline{
						{
							{
								Key: "$geoNear", Value: bson.D{
									{Key: "near", Value: "$$vehicleLoc"},
									{Key: "distanceField", Value: "dist.calculated"},
									{Key: "spherical", Value: true},
								},
							},
						},
						{
							{Key: "$limit", Value: 1},
						},
					}},
					{Key: "as", Value: "nearestParcel"},
				},
			},
		},
		{
			{Key: "$unwind", Value: "$nearestParcel"},
		},
		{
			{
				Key: "$project", Value: bson.D{
					{Key: "number", Value: 1},
					{Key: "vehicleLoc", Value: "$loc"},
					{Key: "nearestParcel", Value: bson.D{
						{Key: "_id", Value: "$nearestParcel._id"},
						{Key: "loc", Value: "$nearestParcel.loc"},
						{Key: "distance", Value: "$nearestParcel.dist.calculated"},
					}},
				},
			},
		},
	}

	cursor, err := r.Collection.Aggregate(r.DB.Context, pipeline)
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
