package main

import "fmt"

type externalRequest struct {
	direction      Direction
	currentFloor   int
	requestedFloor int
}

func (this externalRequest) isExternalRequest() bool {
	return true
}

func (this externalRequest) getRequestedFloor() int {
	return this.requestedFloor
}

func (this externalRequest) getDirection() Direction {
	return this.direction
}

func (this externalRequest) getCurrentFloor() (int, error) {
	return 0, fmt.Errorf("currentFloor is not applicable for internal request")
}
