package main

type ParkingTicketStatus uint8

const (
	active ParkingTicketStatus = iota
	paid
)

func (s ParkingTicketStatus) toString() string {
	switch s {
	case active:
		return "Active"
	case paid:
		return "Paid"
	}
	return ""
}
