package main

import "fmt"

type hasMoneyState struct {
	atmMachine *atmMachine
}

func (i *hasMoneyState) stateName() string {
	return "hasMoneyState"
}

func (i *hasMoneyState) addMoney(money int) error {
	fmt.Errorf("Add money in progress")
	i.atmMachine.incrementMoney(money)
	return nil
}

func (i *hasMoneyState) enterAmount(money int) error {
	fmt.Errorf("Amount is entered. Amount:%n", money)
	err := i.atmMachine.checkAvailability(money)
	if err != nil {
		return err
	}
	i.atmMachine.setState(i.atmMachine.amountEntered)
	return nil
}

func (i *hasMoneyState) enterPin(pin int) error {
	return fmt.Errorf("First enter the amount")
}
func (i *hasMoneyState) dispenseCash(money int) error {
	return fmt.Errorf("First enter the amount")
}
