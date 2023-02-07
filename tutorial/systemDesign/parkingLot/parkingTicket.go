package main

import (
	"fmt"
	"time"
)

type ParkingTicket struct {
	vehicle     Vehicle
	parkingSpot ParkingSpot
	status      ParkingTicketStatus
	entryTime   time.Time
	exitTime    time.Time
	entryGate   Gate
	exitGate    Gate
	price       int
	pgType      PaymentGatewayType
}

func initParkingTicket(vehicle Vehicle, pSpot ParkingSpot, entryGate Gate) *ParkingTicket {
	return &ParkingTicket{
		vehicle:     vehicle,
		parkingSpot: pSpot,
		status:      active,
		entryTime:   time.Now(),
		entryGate:   entryGate,
	}
}

func (this *ParkingTicket) exit(pgType PaymentGatewayType) {
	fmt.Printf("Vehicle with number %s exiting from Parking Lot\n", this.vehicle.numberPlate)
	this.exitTime = time.Now()
	pRateInstance := pRateFactorySingleInstance.getParkingRateInstanceByVehicleType(this.vehicle.vehicleType)
	totalDurationInHours := int(this.exitTime.Sub(this.entryTime).Hours())
	this.price = pRateInstance.amountToPay(totalDurationInHours) + 1
	this.pgType = pgType
	pgInstance := pgFactorySingleInstance.getPaymentGatewayInstanceByPGType(pgType)
	pgInstance.pay(this.price)
	this.status = paid
}

func (this *ParkingTicket) print() {
	fmt.Printf("Issued ticket for vehicle number %s at parking spot %s\n ", this.vehicle.numberPlate, this.parkingSpot.getLocation())
	//fmt.Printf("\nPrinting Ticket\n")
	//fmt.Printf("Status: %s, \nEntryTime: %s, \nEntryGate: %d, \nVehicle: %s, \nParking Spot: \n\n", this.status.toString(), this.entryTime.String(), this.entryGate, this.vehicle.toString())
}
