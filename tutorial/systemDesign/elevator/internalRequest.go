package main

import "fmt"

type internalRequest struct {
	requestedFloor int
	direction      Direction
}

func (this internalRequest) isExternalRequest() bool {
	return false
}

func (this internalRequest) getRequestedFloor() int {
	return this.requestedFloor
}

func (this internalRequest) getDirection() Direction {
	return this.direction
}

func (this internalRequest) getCurrentFloor() (int, error) {
	return 0, fmt.Errorf("currentFloor is not applicable for internal request")

}
