package main

import (
	"fmt"
	"log"
)

func main() {
	atmMachine := newATMMachine(100)

	fmt.Println("<<<<<First Transactin: Withdrawing amount 10>>>> ")

	fmt.Printf("ATM current state %s\n\n", atmMachine.currentState.stateName())
	err := atmMachine.enterAmount(10)
	if err != nil {
		log.Fatalf(err.Error())

	}
	fmt.Printf("Amount Entered: %d\n", 10)
	fmt.Printf("Atm Total Money: %d\n", atmMachine.totalMoney)
	fmt.Printf("ATM current state %s\n\n", atmMachine.currentState.stateName())

	err = atmMachine.enterPin(1234)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Pin Entered: %d\n", 10)
	fmt.Printf("ATM Total Money: %d\n", atmMachine.totalMoney)
	fmt.Printf("ATM current state %s\n\n", atmMachine.currentState.stateName())

	err = atmMachine.dispenseCash(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Dispense Cash: %d\n", 10)
	fmt.Printf("ATM Total Money: %d\n", atmMachine.totalMoney)
	fmt.Printf("ATM current state %s\n\n", atmMachine.currentState.stateName())

	fmt.Println()

	fmt.Println("<<<<<Second Transactin: Admin adding 50>>>>")
	err = atmMachine.addMoney(50)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Atm Total Money: %d\n", atmMachine.totalMoney)
	fmt.Printf("ATM current state %s\n\n", atmMachine.currentState.stateName())

	fmt.Println("<<<<<Third Transactin. Withdrawing amount 200>>>>")
	err = atmMachine.enterAmount(200)
	if err != nil {
		log.Fatalf(err.Error())

	}
}
