package main

import "math"

type shortestTime struct {
}

func (this *shortestTime) processExternalRequest(elevators []*elevator, eRequest externalRequest) error {
	direction := eRequest.direction
	currentFloor := eRequest.currentFloor

	shortestTime := math.MaxInt64
	var selectedElevator *elevator
	for _, elev := range elevators {
		time, err := elev.getTimeToReach(currentFloor, direction)
		if err != nil {
			return err
		}
		if time < shortestTime {
			shortestTime = time
			selectedElevator = elev
		}
	}

	selectedElevator.addExternalRequest(eRequest)
	return nil
}
