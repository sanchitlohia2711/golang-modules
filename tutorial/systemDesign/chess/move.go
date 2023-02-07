package main

type move struct {
	currLocation location
	newLocation  location
	piece        piece
	resign       bool
	drawOffer    bool
}
