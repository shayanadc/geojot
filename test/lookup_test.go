package test

import (
	"geo-jot/service"
	"testing"
)

func TestLookUpService(t *testing.T) {

	parcelA := service.Parcel{
		Longtitude: 1.8,
		Latitude:   1.5,
		Title:      "Parcel A",
	}

	vehicleA := service.Vehicle{
		Longtitude: 0.9,
		Latitude:   0.9,
		Title:      "Vehicle A",
	}

	parcelDistance := vehicleA.DistanceTo(parcelA)

	if parcelDistance != 120.2570602473858 {
		t.Errorf("Distance should be 0, but got %f", parcelDistance)
	}

	vehicleDistance := parcelA.DistanceTo(vehicleA)

	if vehicleDistance != 120.2570602473858 {
		t.Errorf("Distance should be 0, but got %f", vehicleDistance)
	}
}
