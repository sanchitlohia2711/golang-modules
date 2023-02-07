package main

type TruckParkingSpot struct {
	full     bool
	floor    int
	location string
}

func (this *TruckParkingSpot) isFull() bool {
	return this.full
}

func (this *TruckParkingSpot) getFloor() int {
	return this.floor
}

func (this *TruckParkingSpot) getLocation() string {
	return this.location
}

func (this *TruckParkingSpot) getParkingSpotType() ParkingSpotType {
	return truckPT
}

func (this *TruckParkingSpot) markFull() {
	this.full = true
}

func (this *TruckParkingSpot) markFree() {
	this.full = true
}
