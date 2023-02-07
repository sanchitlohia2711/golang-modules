package main

type ParkingSpotType uint8

const (
	carPT ParkingSpotType = iota
	truckPT
	motorcyclePT
)

func (s ParkingSpotType) toString() string {
	switch s {
	case carPT:
		return "Car Parking Type"
	case truckPT:
		return "Truck Parking Type"
	case motorcyclePT:
		return "Motorcylce Parking Type"
	}
	return ""
}

func initParkingSpot(floor int, partkingSpotType ParkingSpotType, location string) ParkingSpot {
	switch partkingSpotType {
	case carPT:
		return &CarParkingSpot{full: false,
			floor:    floor,
			location: location,
		}
	case truckPT:
		return &TruckParkingSpot{full: false,
			floor:    floor,
			location: location,
		}
	case motorcyclePT:
		return &MotorCycleParkingSpot{full: false,
			floor:    floor,
			location: location,
		}
	}
	return nil
}
