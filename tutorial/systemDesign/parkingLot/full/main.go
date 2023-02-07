package main

import (
	"fmt"
	"time"
)

type CarParkingSpot struct {
	full     bool
	floor    int
	location string
}

func (this *CarParkingSpot) isFull() bool {
	return this.full
}

func (this *CarParkingSpot) getFloor() int {
	return this.floor
}

func (this *CarParkingSpot) getLocation() string {
	return this.location
}

func (this *CarParkingSpot) getParkingSpotType() ParkingSpotType {
	return carPT
}

func (this *CarParkingSpot) markFull() {
	this.full = true
}

func (this *CarParkingSpot) markFree() {
	this.full = true
}

type ParkingSpotDLLNode struct {
	pSpot ParkingSpot
	prev  *ParkingSpotDLLNode
	next  *ParkingSpotDLLNode
}

type ParkingSpotDLL struct {
	len  int
	tail *ParkingSpotDLLNode
	head *ParkingSpotDLLNode
}

func initDoublyList() *ParkingSpotDLL {
	return &ParkingSpotDLL{}
}

func (d *ParkingSpotDLL) AddToFront(node *ParkingSpotDLLNode) {

	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.len++
	return
}

func (d *ParkingSpotDLL) RemoveFromFront() {
	if d.head == nil {
		return
	} else if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
	}
	d.len--
}

func (d *ParkingSpotDLL) AddToEnd(node *ParkingSpotDLLNode) {
	newNode := node
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
}
func (d *ParkingSpotDLL) Front() *ParkingSpotDLLNode {
	return d.head
}

func (d *ParkingSpotDLL) MoveNodeToEnd(node *ParkingSpotDLLNode) {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
	if d.tail == node {
		d.tail = prev
	}
	if d.head == node {
		d.head = next
	}
	node.next = nil
	node.prev = nil
	d.len--
	d.AddToEnd(node)
}

func (d *ParkingSpotDLL) MoveNodeToFront(node *ParkingSpotDLLNode) {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
	if d.tail == node {
		d.tail = prev
	}
	if d.head == node {
		d.head = next
	}
	node.next = nil
	node.prev = nil
	d.len--
	d.AddToFront(node)
}

func (d *ParkingSpotDLL) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("Location = %v, parkingType = %s, floor = %d full =%t\n", temp.pSpot.getLocation(), temp.pSpot.getParkingSpotType().toString(), temp.pSpot.getFloor(), temp.pSpot.isFull())
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *ParkingSpotDLL) Size() int {
	return d.len
}

type Gate struct {
	floor    int
	gateType GateType
}

func initGate(floor int, gateType GateType) Gate {
	return Gate{
		floor:    floor,
		gateType: gateType,
	}
}

type GateType uint8

const (
	entryGateType GateType = iota
	exitGateType  GateType = iota
)

var (
	pRateFactorySingleInstance = &ParkingRateFactory{
		parkingRateMap: make(map[VehicleType]ParkingRate),
	}
	pgFactorySingleInstance = &PaymentGatewayFactory{
		paymentGatewayMap: make(map[PaymentGatewayType]PaymentGateway),
	}
)

type MotorCycleParkingSpot struct {
	full     bool
	floor    int
	location string
}

func (this *MotorCycleParkingSpot) isFull() bool {
	return this.full
}

func (this *MotorCycleParkingSpot) getFloor() int {
	return this.floor
}

func (this *MotorCycleParkingSpot) getLocation() string {
	return this.location
}

func (this *MotorCycleParkingSpot) getParkingSpotType() ParkingSpotType {
	return motorcyclePT
}

func (this *MotorCycleParkingSpot) markFull() {
	this.full = true
}

func (this *MotorCycleParkingSpot) markFree() {
	this.full = true
}

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

type ParkingRateFactory struct {
	parkingRateMap map[VehicleType]ParkingRate
}

func (this *ParkingRateFactory) getParkingRateInstanceByVehicleType(vType VehicleType) ParkingRate {
	if this.parkingRateMap[vType] != nil {
		return this.parkingRateMap[vType]
	}
	if vType == car {
		this.parkingRateMap[vType] = &CarParkingRate{
			ParkingRateBase{perHourCharges: 2},
		}
		return this.parkingRateMap[vType]
	}
	if vType == truck {
		this.parkingRateMap[vType] = &TruckParkingRate{
			ParkingRateBase{perHourCharges: 3},
		}
		return this.parkingRateMap[vType]
	}
	if vType == motorcycle {
		this.parkingRateMap[vType] = &MotorCycleParkingRate{
			ParkingRateBase{perHourCharges: 1},
		}
		return this.parkingRateMap[vType]
	}
	return nil
}

type ParkingRateBase struct {
	perHourCharges int
}

type ParkingRate interface {
	amountToPay(int) int
}

type CarParkingRate struct {
	ParkingRateBase
}

func (this *CarParkingRate) amountToPay(hours int) int {
	return this.perHourCharges * hours
}

type TruckParkingRate struct {
	ParkingRateBase
}

func (this *TruckParkingRate) amountToPay(hours int) int {
	return this.perHourCharges * hours
}

type MotorCycleParkingRate struct {
	ParkingRateBase
}

func (this *MotorCycleParkingRate) amountToPay(hours int) int {
	return this.perHourCharges * hours
}

type ParkingSpot interface {
	isFull() bool
	getFloor() int
	getLocation() string
	getParkingSpotType() ParkingSpotType
	markFull()
	markFree()
}

type ParkingSpotType uint8

const (
	carPT ParkingSpotType = iota
	truckPT
	motorcyclePT
)

func (s ParkingSpotType) toString() string {
	switch s {
	case carPT:
		return "Car Parking Type"
	case truckPT:
		return "Truck Parking Type"
	case motorcyclePT:
		return "Motorcylce Parking Type"
	}
	return ""
}

func initParkingSpot(floor int, partkingSpotType ParkingSpotType, location string) ParkingSpot {
	switch partkingSpotType {
	case carPT:
		return &CarParkingSpot{full: false,
			floor:    floor,
			location: location,
		}
	case truckPT:
		return &TruckParkingSpot{full: false,
			floor:    floor,
			location: location,
		}
	case motorcyclePT:
		return &MotorCycleParkingSpot{full: false,
			floor:    floor,
			location: location,
		}
	}
	return nil
}

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

type ParkingTicketStatus uint8

const (
	active ParkingTicketStatus = iota
	paid
)

func (s ParkingTicketStatus) toString() string {
	switch s {
	case active:
		return "Active"
	case paid:
		return "Paid"
	}
	return ""
}

type PaymentGatewayType uint8

const (
	cash PaymentGatewayType = iota
	creditCard
	debitCard
)

type PaymentGatewayFactory struct {
	paymentGatewayMap map[PaymentGatewayType]PaymentGateway
}

func (this *PaymentGatewayFactory) getPaymentGatewayInstanceByPGType(pgType PaymentGatewayType) PaymentGateway {
	if this.paymentGatewayMap[pgType] != nil {
		return this.paymentGatewayMap[pgType]
	}
	if pgType == cash {
		this.paymentGatewayMap[pgType] = &CashPaymentGateway{}
		return this.paymentGatewayMap[pgType]
	}
	if pgType == creditCard {
		this.paymentGatewayMap[pgType] = &CreditCardPaymentGateway{}
		return this.paymentGatewayMap[pgType]
	}
	if pgType == debitCard {
		this.paymentGatewayMap[pgType] = &DebitCardPaymentGateway{}
		return this.paymentGatewayMap[pgType]
	}
	return nil
}

type PaymentGateway interface {
	pay(int)
}

type CashPaymentGateway struct {
}

func (this CashPaymentGateway) pay(price int) {
	fmt.Printf("Paying price of %d$ through cash payment\n", price)
}

type CreditCardPaymentGateway struct {
}

func (this CreditCardPaymentGateway) pay(price int) {
	fmt.Printf("Paying price of %d$ through credit card payment\n", price)
}

type DebitCardPaymentGateway struct {
}

func (this DebitCardPaymentGateway) pay(price int) {
	fmt.Printf("Paying price of %d$ through debit card payment\n", price)
}

type TruckParkingSpot struct {
	full     bool
	floor    int
	location string
}

func (this *TruckParkingSpot) isFull() bool {
	return this.full
}

func (this *TruckParkingSpot) getFloor() int {
	return this.floor
}

func (this *TruckParkingSpot) getLocation() string {
	return this.location
}

func (this *TruckParkingSpot) getParkingSpotType() ParkingSpotType {
	return truckPT
}

func (this *TruckParkingSpot) markFull() {
	this.full = true
}

func (this *TruckParkingSpot) markFree() {
	this.full = true
}

type Vehicle struct {
	numberPlate string
	vehicleType VehicleType
}

func (v Vehicle) toString() string {
	return fmt.Sprintf("{NumberPlate: %s, VehicleType: %s}", v.numberPlate, v.vehicleType.toString())
}

type VehicleType uint8

const (
	car VehicleType = iota
	truck
	motorcycle
)

func (s VehicleType) toString() string {
	switch s {
	case car:
		return "Car"
	case truck:
		return "Truck"
	case motorcycle:
		return "Motorcylce"
	}
	return ""
}

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
