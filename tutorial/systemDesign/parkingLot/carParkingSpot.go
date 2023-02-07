package main

type CarParkingSpot struct {
	full     bool
	floor    int
	location string
}

func (this *CarParkingSpot) isFull() bool {
	return this.full
}

func (this *CarParkingSpot) getFloor() int {
	return this.floor
}

func (this *CarParkingSpot) getLocation() string {
	return this.location
}

func (this *CarParkingSpot) getParkingSpotType() ParkingSpotType {
	return carPT
}

func (this *CarParkingSpot) markFull() {
	this.full = true
}

func (this *CarParkingSpot) markFree() {
	this.full = true
}
