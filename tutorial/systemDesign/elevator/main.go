package main

func main() {
	elevator1 := initializeElevator(1, 10, -1)
	elevator2 := initializeElevator(2, 10, -1)

	elevAlgo := shortestTime{}

	elevSystem := elevatorSystem{
		elevators: []*elevator{elevator1, elevator2},
		elevAlgo:  elevAlgo,
	}

	eReq1 := externalRequest{
		currentFloor:   0,
		direction:      upDirection,
		requestedFloor: 1,
	}

	elevSystem.processExternalRequest(eReq1)

}
