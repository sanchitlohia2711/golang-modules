package main

import "fmt"

type amountEnteredState struct {
	atmMachine *atmMachine
}

func (i *amountEnteredState) stateName() string {
	return "amountEnteredState"
}

func (i *amountEnteredState) addMoney(money int) error {
	return fmt.Errorf("Dispensing process in progress")
}

func (i *amountEnteredState) enterAmount(money int) error {
	return fmt.Errorf("Amount already entered")
}

func (i *amountEnteredState) enterPin(pin int) error {
	err := i.atmMachine.verifyPin(pin)
	if err != nil {
		return err
	}
	i.atmMachine.setState(i.atmMachine.pinEntered)
	return nil
}
func (i *amountEnteredState) dispenseCash(money int) error {
	return fmt.Errorf("First Enter Pin")
}
