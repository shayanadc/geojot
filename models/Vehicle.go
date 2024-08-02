package models

type Vehicle struct {
	Number      string    `bson:"number"`
	Coordinates []float64 `bson:"coordinates"`
}
