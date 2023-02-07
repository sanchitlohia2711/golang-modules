package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type elevator struct {
	mu               *sync.Mutex
	id               int
	direction        Direction
	state            ElevatorState
	currentFloor     int
	destFloorUp      int
	destFloorUpSet   bool
	destFloorDown    int
	destFloorDownSet bool
	floorMap         map[int][]request
	maxFloor         int
	minFloor         int
}

func initializeElevator(id int, maxFloor int, minFloor int) *elevator {
	mutex := &sync.Mutex{}
	return &elevator{
		mu:       mutex,
		id:       id,
		maxFloor: maxFloor,
		minFloor: minFloor,
	}
}

func (this *elevator) handleAllRequestsOnCurrentFloor() {
	requests, ok := this.floorMap[this.currentFloor]
	if ok {
		fmt.Printf("Moving Up: Opening elevator with Id:%d on Floor:%d. Requested Floor: %d", this.id, this.currentFloor, req.requestedFloor)
	}
	for _, req := range requests {
		if req.isExternalRequest() {
			internalRequest := internalRequest{
				requestedFloor: req.getRequestedFloor(),
				direction:      req.getDirection(),
			}
			this.floorMap[req.getRequestedFloor()] = append(this.floorMap[req.getRequestedFloor()], internalRequest)
			if this.movingUp() {
				if this.destFloorUpSet {
					if this.destFloorUp < req.getRequestedFloor() {
						this.destFloorUp = req.getRequestedFloor()
					}
				} else {
					this.destFloorUpSet = true
					this.destFloorUp = req.getRequestedFloor()
				}
			} else {
				if this.destFloorDownSet {
					if this.destFloorDown > req.getRequestedFloor() {
						this.destFloorDown = req.getRequestedFloor()
					}
				} else {
					this.destFloorDownSet = true
					this.destFloorDown = req.getRequestedFloor()
				}
			}
		}
	}

	delete(this.floorMap, this.currentFloor)
}

func (this *elevator) start() {
	fmt.Printf("Starting elevator with Id:%d", this.id)
	for {
		time.Sleep(time.Second * 1)
		this.mu.Lock()
		if this.isIdle() {
			continue
		}

		if this.movingUp() {
			this.currentFloor++
			this.handleAllRequestsOnCurrentFloor()
			if this.currentFloor == this.destFloorUp {
				this.destFloorUpSet = false
				if this.destFloorDownSet {
					this.changeDirectionToDown()
				} else {
					this.moveToIdle()
				}
			}
		}

		if this.movingUp() {
			this.currentFloor--

			delete(this.floorMap, this.currentFloor)
			if this.currentFloor == this.destFloorDown {
				this.destFloorDownSet = false
				if this.destFloorUpSet {
					this.changeDirectionToUp()
				} else {
					this.moveToIdle()
				}
			}
		}
		this.mu.Unlock()
	}
}

func (this *elevator) changeDirectionToDown() {
	this.direction = downDirection
}

func (this *elevator) changeDirectionToUp() {
	this.direction = upDirection
}

func (this *elevator) moveToIdle() {
	this.state = ElevatorIdle
}

func (this *elevator) getTimeToReach(requestedFloor int, requestedDirection Direction) (int, error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if this.isIdle() {
		return int(math.Abs(float64(requestedFloor - this.currentFloor))), nil
	}
	if this.inDirectionOfRequest(requestedFloor, requestedDirection) {
		return int(math.Abs(float64(requestedFloor - this.currentFloor))), nil
	}

	if this.oppDirectionOfRequest(requestedFloor, requestedDirection) {
		if this.direction == upDirection {
			distanceTillDestFloorUp := 0
			distanceFromDestFloorToReqFloor := 0
			if this.destFloorUpSet {
				distanceTillDestFloorUp = int(math.Abs(float64(this.destFloorUp - this.currentFloor)))
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(this.currentFloor - requestedFloor)))
			} else {
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(distanceTillDestFloorUp - requestedFloor)))

			}
			return distanceTillDestFloorUp + distanceFromDestFloorToReqFloor, nil
		} else {

			distanceTillDestFloorDown := 0
			distanceFromDestFloorToReqFloor := 0
			if this.destFloorDownSet {
				distanceTillDestFloorDown = int(math.Abs(float64(this.destFloorDown - this.currentFloor)))
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(this.destFloorDown - requestedFloor)))
			} else {
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(this.currentFloor - requestedFloor)))

			}
			return distanceTillDestFloorDown + distanceFromDestFloorToReqFloor, nil
		}

	}

	if this.awayDirectionOfRequest(requestedFloor, requestedDirection) {
		if this.direction == upDirection {
			distanceTillDestFloorUp := 0
			distanceFromDestFloorToReqFloor := 0
			if this.destFloorUpSet {
				distanceTillDestFloorUp = int(math.Abs(float64(this.destFloorUp - this.currentFloor)))
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(this.currentFloor - requestedFloor)))
			} else {
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(distanceTillDestFloorUp - requestedFloor)))

			}
			return distanceTillDestFloorUp + distanceFromDestFloorToReqFloor, nil
		} else {
			distanceTillDestFloorDown := 0
			distanceFromDestFloorToReqFloor := 0
			if this.destFloorDownSet {
				distanceTillDestFloorDown = int(math.Abs(float64(this.destFloorDown - this.currentFloor)))
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(this.destFloorDown - requestedFloor)))
			} else {
				distanceFromDestFloorToReqFloor = int(math.Abs(float64(this.currentFloor - requestedFloor)))

			}
			return distanceTillDestFloorDown + distanceFromDestFloorToReqFloor, nil
		}
	}

	return 0, fmt.Errorf("Wrong elevator State")

}

func (this *elevator) isIdle() bool {
	return this.state == ElevatorIdle
}

func (this *elevator) movingUp() bool {
	return this.direction == upDirection
}

func (this *elevator) movingDown() bool {
	return this.direction == downDirection
}

func (this *elevator) inDirectionOfRequest(requestedFloor int, requestedDirection Direction) bool {
	if this.direction != requestedDirection {
		return false
	}

	if this.direction == upDirection && this.currentFloor < requestedFloor {
		return true
	}

	if this.direction == downDirection && this.currentFloor > requestedFloor {
		return true
	}

	return false
}

func (this *elevator) oppDirectionOfRequest(requestedFloor int, requestedDirection Direction) bool {
	if this.direction != requestedDirection {
		return true
	}

	return false
}

func (this *elevator) addExternalRequest(eReq externalRequest) {
	this.floorMap[eReq.getRequestedFloor()] = append(this.floorMap[eReq.getRequestedFloor()], eReq)
}

func (this *elevator) awayDirectionOfRequest(requestedFloor int, requestedDirection Direction) bool {
	if this.direction != requestedDirection {
		return false
	}

	if this.direction == upDirection && this.currentFloor >= requestedFloor {
		return true
	}

	if this.direction == downDirection && this.currentFloor <= requestedFloor {
		return true
	}

	return false
}
