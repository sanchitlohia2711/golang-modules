package main

import "fmt"

type ParkingSystem struct {
	issuedTickets map[string]ParkingTicket
	parkingFloor  []*ParkingFloor
	isFull        map[ParkingSpotType]bool
	entryGates    []Gate
	exitGates     []Gate
}

func (this *ParkingSystem) addParkingSpot(parkingSpot ParkingSpot) {
	this.parkingFloor[parkingSpot.getFloor()-1].addParkingSpot(parkingSpot)
}

func (this *ParkingSystem) bookParkingSpot(pSpotType ParkingSpotType) (ParkingSpot, error) {
	for _, pFloor := range this.parkingFloor {
		pSpot, err := pFloor.bookParkingSpot(pSpotType)
		if err == nil {
			return pSpot, nil
		}
	}

	return nil, fmt.Errorf("Cannot issue ticket. All %s parking spot type are full\n", pSpotType.toString())
}

func (this *ParkingSystem) issueTicket(pSpotType ParkingSpotType, vehicle Vehicle, entryGate Gate) (*ParkingTicket, error) {
	fmt.Printf("\nGoing to issue ticket for vehicle number %s\n", vehicle.numberPlate)
	pSpot, err := this.bookParkingSpot(pSpotType)
	if err != nil {
		return nil, err
	}

	ticket := initParkingTicket(vehicle, pSpot, entryGate)
	return ticket, nil
}

func (this *ParkingSystem) printStatus() {
	fmt.Println("\nPrinting Status of Parking Spot")
	for _, pFloor := range this.parkingFloor {
		pFloor.printStatus()
	}
}

func (this *ParkingSystem) exitVehicle(ticket *ParkingTicket, pGType PaymentGatewayType) {
	this.parkingFloor[ticket.parkingSpot.getFloor()-1].freeParkingSpot(ticket.parkingSpot)
	ticket.exit(pGType)
}
