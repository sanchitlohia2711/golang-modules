package main

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
