package main

type MotorCycleParkingSpot struct {
	full     bool
	floor    int
	location string
}

func (this *MotorCycleParkingSpot) isFull() bool {
	return this.full
}

func (this *MotorCycleParkingSpot) getFloor() int {
	return this.floor
}

func (this *MotorCycleParkingSpot) getLocation() string {
	return this.location
}

func (this *MotorCycleParkingSpot) getParkingSpotType() ParkingSpotType {
	return motorcyclePT
}

func (this *MotorCycleParkingSpot) markFull() {
	this.full = true
}

func (this *MotorCycleParkingSpot) markFree() {
	this.full = true
}
