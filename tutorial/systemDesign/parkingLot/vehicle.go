package main

import "fmt"

type Vehicle struct {
	numberPlate string
	vehicleType VehicleType
}

func (v Vehicle) toString() string {
	return fmt.Sprintf("{NumberPlate: %s, VehicleType: %s}", v.numberPlate, v.vehicleType.toString())
}
