package main

import (
	"fmt"
)

var (
	pRateFactorySingleInstance = &ParkingRateFactory{
		parkingRateMap: make(map[VehicleType]ParkingRate),
	}
	pgFactorySingleInstance = &PaymentGatewayFactory{
		paymentGatewayMap: make(map[PaymentGatewayType]PaymentGateway),
	}
)

func main() {

	testSmallParkingLot()

	//testLargeParkingLot()

}

func testSmallParkingLot() {
	firstParkingFloor := initParkingFloor(1)
	firstFloorEntryGate1 := initGate(1, entryGateType)
	firstFloorExitGate := initGate(1, exitGateType)
	firstFloorCarParkingSpot1 := initParkingSpot(1, carPT, "A1")
	firstFloorCarParkingSpot2 := initParkingSpot(1, carPT, "A2")

	parkingSystem := ParkingSystem{
		parkingFloor:  []*ParkingFloor{firstParkingFloor},
		entryGates:    []Gate{firstFloorEntryGate1},
		exitGates:     []Gate{firstFloorExitGate},
		issuedTickets: make(map[string]ParkingTicket),
	}
	//Add first floor parkings
	parkingSystem.addParkingSpot(firstFloorCarParkingSpot1)
	parkingSystem.addParkingSpot(firstFloorCarParkingSpot2)

	carVehicle1 := Vehicle{
		numberPlate: "C1",
		vehicleType: car,
	}
	carVehicle2 := Vehicle{
		numberPlate: "C2",
		vehicleType: car,
	}

	parkingSystem.printStatus()
	carVehicleTicket1, err := parkingSystem.issueTicket(carPT, carVehicle1, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket1.print()

	carVehicleTicket2, err := parkingSystem.issueTicket(carPT, carVehicle2, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket2.print()

	carVehicle3 := Vehicle{
		numberPlate: "C3",
		vehicleType: car,
	}
	carVehicleTicket3, err := parkingSystem.issueTicket(carPT, carVehicle3, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	parkingSystem.printStatus()

	parkingSystem.exitVehicle(carVehicleTicket1, cash)
	parkingSystem.printStatus()

	carVehicleTicket3, err = parkingSystem.issueTicket(carPT, carVehicle3, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket3.print()
	parkingSystem.printStatus()

}

func testLargeParkingLot() {
	//We have two parking floor
	firstParkingFloor := initParkingFloor(1)
	secondParkingFloor := initParkingFloor(2)

	//We have two entry gates in firstParkingFloor
	firstFloorEntryGate1 := initGate(1, entryGateType)
	firstFloorEntryGate2 := initGate(1, entryGateType)

	//We have one exit gate on firstParkingFloor
	firstFloorExitGate := initGate(1, exitGateType)

	parkingSystem := ParkingSystem{
		parkingFloor:  []*ParkingFloor{firstParkingFloor, secondParkingFloor},
		entryGates:    []Gate{firstFloorEntryGate1, firstFloorEntryGate2},
		exitGates:     []Gate{firstFloorExitGate},
		issuedTickets: make(map[string]ParkingTicket),
	}

	//We have two car parking spots, two motorcyle parking spots, 1 truck paring spot on each of the floor
	firstFloorCarParkingSpot1 := initParkingSpot(1, carPT, "A1")
	firstFloorCarParkingSpot2 := initParkingSpot(1, carPT, "A2")
	firstFloorMotorCycleParkingSpot1 := initParkingSpot(1, motorcyclePT, "A3")
	firstFloorMotorCycleParkingSpot2 := initParkingSpot(1, motorcyclePT, "A4")
	firstFloorTruckParkingSpot := initParkingSpot(1, truckPT, "A5")

	//We have two car parking spots, two motorcyle parking spots, 1 truck paring spot on each of the floor
	secondFloorCarParkingSpot1 := initParkingSpot(2, carPT, "B1")
	secondFloorCarParkingSpot2 := initParkingSpot(2, carPT, "B2")
	secondFloorMotorCycleParkingSpot1 := initParkingSpot(2, motorcyclePT, "B3")
	secondFloorMotorCycleParkingSpot2 := initParkingSpot(2, motorcyclePT, "B4")
	secondFloorTruckParkingSpot := initParkingSpot(2, truckPT, "B5")

	//Add first floor parkings
	parkingSystem.addParkingSpot(firstFloorCarParkingSpot1)
	parkingSystem.addParkingSpot(firstFloorCarParkingSpot2)
	parkingSystem.addParkingSpot(firstFloorMotorCycleParkingSpot1)
	parkingSystem.addParkingSpot(firstFloorMotorCycleParkingSpot2)
	parkingSystem.addParkingSpot(firstFloorTruckParkingSpot)

	//Add second floor parkings
	parkingSystem.addParkingSpot(secondFloorCarParkingSpot1)
	parkingSystem.addParkingSpot(secondFloorCarParkingSpot2)
	parkingSystem.addParkingSpot(secondFloorMotorCycleParkingSpot1)
	parkingSystem.addParkingSpot(secondFloorMotorCycleParkingSpot2)
	parkingSystem.addParkingSpot(secondFloorTruckParkingSpot)

	carVehicle1 := Vehicle{
		numberPlate: "C1",
		vehicleType: car,
	}
	carVehicle2 := Vehicle{
		numberPlate: "C2",
		vehicleType: car,
	}
	motorCycleVehicle1 := Vehicle{
		numberPlate: "M1",
		vehicleType: motorcycle,
	}
	motorCycleVehicle2 := Vehicle{
		numberPlate: "M2",
		vehicleType: motorcycle,
	}

	truckVehicle1 := Vehicle{
		numberPlate: "T1",
		vehicleType: motorcycle,
	}

	parkingSystem.printStatus()
	carVehicleTicket1, err := parkingSystem.issueTicket(carPT, carVehicle1, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket1.print()
	carVehicleTicket2, err := parkingSystem.issueTicket(carPT, carVehicle2, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket2.print()
	motorCycleVehicleTicket1, err := parkingSystem.issueTicket(motorcyclePT, motorCycleVehicle1, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	motorCycleVehicleTicket1.print()
	motorCycleVehicleTicket2, err := parkingSystem.issueTicket(motorcyclePT, motorCycleVehicle2, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	motorCycleVehicleTicket2.print()
	truckVehicleTicket1, err := parkingSystem.issueTicket(truckPT, truckVehicle1, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	truckVehicleTicket1.print()
	parkingSystem.printStatus()

	carVehicle3 := Vehicle{
		numberPlate: "C3",
		vehicleType: car,
	}
	carVehicle4 := Vehicle{
		numberPlate: "C4",
		vehicleType: car,
	}
	motorCycleVehicle3 := Vehicle{
		numberPlate: "M3",
		vehicleType: motorcycle,
	}
	motorCycleVehicle4 := Vehicle{
		numberPlate: "M4",
		vehicleType: motorcycle,
	}

	truckVehicle2 := Vehicle{
		numberPlate: "T2",
		vehicleType: motorcycle,
	}
	carVehicleTicket3, err := parkingSystem.issueTicket(carPT, carVehicle3, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket3.print()
	carVehicleTicket4, err := parkingSystem.issueTicket(carPT, carVehicle4, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket4.print()
	motorCycleVehicleTicket3, err := parkingSystem.issueTicket(motorcyclePT, motorCycleVehicle3, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	motorCycleVehicleTicket3.print()
	motorCycleVehicleTicket4, err := parkingSystem.issueTicket(motorcyclePT, motorCycleVehicle4, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	motorCycleVehicleTicket4.print()
	truckVehicleTicket2, err := parkingSystem.issueTicket(truckPT, truckVehicle2, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	truckVehicleTicket2.print()
	parkingSystem.printStatus()

	carVehicle5 := Vehicle{
		numberPlate: "C5",
		vehicleType: car,
	}
	carVehicleTicket5, err := parkingSystem.issueTicket(carPT, carVehicle5, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}

	parkingSystem.printStatus()

	parkingSystem.exitVehicle(carVehicleTicket1, cash)
	parkingSystem.printStatus()

	carVehicleTicket5, err = parkingSystem.issueTicket(carPT, carVehicle5, firstFloorEntryGate1)
	if err != nil {
		fmt.Println(err)
	}
	carVehicleTicket5.print()
	parkingSystem.printStatus()
}
