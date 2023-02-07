package main

import "fmt"

type atmMachine struct {
	hasMoney      state
	noMoney       state
	amountEntered state
	pinEntered    state

	currentState state

	totalMoney int
}

func newATMMachine(totalMoney int) *atmMachine {
	a := &atmMachine{
		totalMoney: totalMoney,
	}
	hasMoneyState := &hasMoneyState{
		atmMachine: a,
	}
	noMoneyState := &noMoneyState{
		atmMachine: a,
	}
	amountEnteredState := &amountEnteredState{
		atmMachine: a,
	}
	pinEnteredState := &pinEnteredState{
		atmMachine: a,
	}

	a.setState(hasMoneyState)
	a.hasMoney = hasMoneyState
	a.noMoney = noMoneyState
	a.amountEntered = amountEnteredState
	a.pinEntered = pinEnteredState
	return a
}

func (v *atmMachine) addMoney(money int) error {
	return v.currentState.addMoney(money)
}

func (v *atmMachine) enterAmount(money int) error {
	return v.currentState.enterAmount(money)
}

func (v *atmMachine) enterPin(money int) error {
	return v.currentState.enterPin(money)
}
func (v *atmMachine) dispenseCash(money int) error {
	return v.currentState.dispenseCash(money)
}

func (v *atmMachine) setState(s state) {
	v.currentState = s
}
func (v *atmMachine) incrementMoney(money int) {
	fmt.Printf("Adding %d money:\n", money)
	v.totalMoney = v.totalMoney + money
}

func (v *atmMachine) decrementMoney(money int) {
	fmt.Printf("Dispensing %d cash:\n", money)
	v.totalMoney = v.totalMoney - money
}

func (v *atmMachine) checkAvailability(money int) error {
	fmt.Printf("Checking Availability\n")
	if money < v.totalMoney {
		return nil
	}
	return fmt.Errorf("Not enough money")
}

func (v *atmMachine) verifyPin(pin int) error {
	fmt.Println("Verifying Pin")
	//Pin is always true
	return nil
}
