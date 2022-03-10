package main

import "fmt"

type noMoneyState struct {
	atmMachine *atmMachine
}

func (i *noMoneyState) stateName() string {
	return "noMoneyState"
}

func (i *noMoneyState) addMoney(money int) error {
	fmt.Errorf("Add money in progress")
	i.atmMachine.incrementMoney(money)
	return nil
}

func (i *noMoneyState) enterAmount(money int) error {
	return fmt.Errorf("Add money first")
}

func (i *noMoneyState) enterPin(pin int) error {
	return fmt.Errorf("Add money first")
}
func (i *noMoneyState) dispenseCash(money int) error {
	return fmt.Errorf("Add money first")
}
