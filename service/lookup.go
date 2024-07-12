package service

import (
	"math"
)

type Parcel struct {
	Longtitude float64
	Latitude   float64
	Title      string
}

type Vehicle struct {
	Longtitude float64
	Latitude   float64
	Title      string
}
type ParcelList struct {
	ParcelList []Parcel
}

type VehicleList struct {
	VehicleList []Vehicle
}

const earthRadiusKm = 6371.0

func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

type Location struct {
	Latitude  float64
	Longitude float64
}

func haversine(coord1, coord2 Location) float64 {

	lat1 := degreesToRadians(coord1.Latitude)
	lon1 := degreesToRadians(coord1.Longitude)
	lat2 := degreesToRadians(coord2.Latitude)
	lon2 := degreesToRadians(coord2.Longitude)

	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadiusKm * c

	return distance
}

func (parcel Parcel) DistanceTo(vehicle Vehicle) float64 {
	parcelCoord := Location{Latitude: parcel.Latitude, Longitude: parcel.Longtitude}
	vehicleCoord := Location{Latitude: vehicle.Latitude, Longitude: vehicle.Longtitude}

	return haversine(parcelCoord, vehicleCoord)
}

func (vehicle Vehicle) DistanceTo(parcel Parcel) float64 {
	parcelCoord := Location{Latitude: parcel.Latitude, Longitude: parcel.Longtitude}
	vehicleCoord := Location{Latitude: vehicle.Latitude, Longitude: vehicle.Longtitude}

	return haversine(vehicleCoord, parcelCoord)
}
