package main

type request interface {
	isExternalRequest() bool
	getRequestedFloor() int
	getDirection() Direction
	getCurrentFloor() (int, error)
}
