package models

type VehicleWithNearestParcel struct {
	Number        string   `bson:"number"`
	VehicleLoc    Location `bson:"vehicleLoc"`
	NearestParcel struct {
		ID       string   `bson:"_id"`
		Loc      Location `bson:"loc"`
		Distance float64  `bson:"distance"`
	} `bson:"nearestParcel"`
}
