package main

type VehicleType uint8

const (
	car VehicleType = iota
	truck
	motorcycle
)

func (s VehicleType) toString() string {
	switch s {
	case car:
		return "Car"
	case truck:
		return "Truck"
	case motorcycle:
		return "Motorcylce"
	}
	return ""
}
