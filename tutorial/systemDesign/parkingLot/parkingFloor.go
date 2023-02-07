package main

import "fmt"

type ParkingFloor struct {
	parkingSpots    map[ParkingSpotType]*ParkingSpotDLL
	parkingSpotsMap map[string]*ParkingSpotDLLNode
	isFull          map[ParkingSpotType]bool
	floor           int
}

func initParkingFloor(floor int) *ParkingFloor {
	return &ParkingFloor{
		floor:           floor,
		parkingSpots:    make(map[ParkingSpotType]*ParkingSpotDLL),
		parkingSpotsMap: make(map[string]*ParkingSpotDLLNode),
		isFull:          make(map[ParkingSpotType]bool),
	}
}

func (this *ParkingFloor) addParkingSpot(pSpot ParkingSpot) {
	dll, ok := this.parkingSpots[pSpot.getParkingSpotType()]
	if ok {
		newNode := &ParkingSpotDLLNode{
			pSpot: pSpot,
		}
		dll.AddToFront(newNode)
		this.parkingSpotsMap[pSpot.getLocation()] = newNode
		return
	}

	dll = &ParkingSpotDLL{}
	this.parkingSpots[pSpot.getParkingSpotType()] = dll
	newNode := &ParkingSpotDLLNode{
		pSpot: pSpot,
	}
	this.parkingSpots[pSpot.getParkingSpotType()].AddToFront(newNode)
	this.parkingSpotsMap[pSpot.getLocation()] = newNode
}

func (this *ParkingFloor) bookParkingSpot(pSpotType ParkingSpotType) (ParkingSpot, error) {
	if this.isFull[pSpotType] {
		return nil, fmt.Errorf("%s Parking Spot is full on %d floor", pSpotType.toString(), this.floor)
	}

	nextPSpot := this.parkingSpots[pSpotType].Front()
	nextPSpot.pSpot.markFull()
	this.parkingSpots[pSpotType].MoveNodeToEnd(nextPSpot)
	if this.parkingSpots[pSpotType].Front().pSpot.isFull() {
		this.isFull[pSpotType] = true
	}
	return nextPSpot.pSpot, nil
}

func (this *ParkingFloor) printStatus() {
	for pSpotType, dll := range this.parkingSpots {
		fmt.Printf("Details of parking spots of type %s on floor %d\n", pSpotType.toString(), this.floor)
		dll.TraverseForward()
	}
}

func (this *ParkingFloor) freeParkingSpot(pSpot ParkingSpot) {
	node := this.parkingSpotsMap[pSpot.getLocation()]
	node.pSpot.markFree()
	this.isFull[pSpot.getParkingSpotType()] = false
	this.parkingSpots[pSpot.getParkingSpotType()].MoveNodeToFront(node)
}
