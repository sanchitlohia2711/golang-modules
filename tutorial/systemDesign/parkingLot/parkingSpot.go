package main

type ParkingSpot interface {
	isFull() bool
	getFloor() int
	getLocation() string
	getParkingSpotType() ParkingSpotType
	markFull()
	markFree()
}
