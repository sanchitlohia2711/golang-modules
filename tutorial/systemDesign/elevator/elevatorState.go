package main

type ElevatorState uint8

const (
	ElevatorMoving ElevatorState = iota
	ElevatorIdle
	ElevatorNotWorking
)
