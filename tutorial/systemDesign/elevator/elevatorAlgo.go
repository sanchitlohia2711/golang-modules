package main

type elevatorAlgo interface {
	processExternalRequest([]elevator, externalRequest)
	processInternalRequest(elevator, internalRequest)
}
