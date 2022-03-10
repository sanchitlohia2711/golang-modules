package main

import "fmt"

type pinEnteredState struct {
	atmMachine *atmMachine
}

func (i *pinEnteredState) stateName() string {
	return "pinEnteredState"
}

func (i *pinEnteredState) addMoney(money int) error {
	return fmt.Errorf("Dispensing process in progress")
}

func (i *pinEnteredState) enterAmount(money int) error {
	return fmt.Errorf("Amount and pin already entered")
}

func (i *pinEnteredState) enterPin(pin int) error {
	return fmt.Errorf("Pin already entered")
}
func (i *pinEnteredState) dispenseCash(money int) error {
	i.atmMachine.decrementMoney(money)
	i.atmMachine.setState(i.atmMachine.hasMoney)
	return nil
}
