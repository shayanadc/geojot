package repository

import (
	"geo-jot/container"
	"geo-jot/db"
	"geo-jot/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const vehicleCollection string = "vehicles"

type vehicleRepository struct {
	db         *db.Client
	collection *mongo.Collection
}

func NewVehicleRepository() *vehicleRepository {
	db := container.GetContainer().GetDBClient()

	return &vehicleRepository{db: db, collection: db.GetCollection(vehicleCollection)}
}

func (r *vehicleRepository) GetVehiclesWithNearestParcel() ([]models.VehicleWithNearestParcel, error) {

	cursor, err := r.collection.Aggregate(r.db.Context, bson.A{
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
	defer cursor.Close(r.db.Context)

	var results []models.VehicleWithNearestParcel
	if err = cursor.All(r.db.Context, &results); err != nil {
		log.Fatal(err)
	}
	return results, nil
}
