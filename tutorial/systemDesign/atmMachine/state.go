package main

type state interface {
	stateName() string
	addMoney(int) error
	enterAmount(int) error
	enterPin(int) error
	dispenseCash(int) error
}
