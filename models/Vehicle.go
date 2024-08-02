package models

type Vehicle struct {
	Number      string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}
