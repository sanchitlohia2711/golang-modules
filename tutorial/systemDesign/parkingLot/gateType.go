package main

type GateType uint8

const (
	entryGateType GateType = iota
	exitGateType  GateType = iota
)
