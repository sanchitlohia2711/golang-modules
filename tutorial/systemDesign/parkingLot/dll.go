package main

import "fmt"

type ParkingSpotDLLNode struct {
	pSpot ParkingSpot
	prev  *ParkingSpotDLLNode
	next  *ParkingSpotDLLNode
}

type ParkingSpotDLL struct {
	len  int
	tail *ParkingSpotDLLNode
	head *ParkingSpotDLLNode
}

func initDoublyList() *ParkingSpotDLL {
	return &ParkingSpotDLL{}
}

func (d *ParkingSpotDLL) AddToFront(node *ParkingSpotDLLNode) {

	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.len++
	return
}

func (d *ParkingSpotDLL) RemoveFromFront() {
	if d.head == nil {
		return
	} else if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
	}
	d.len--
}

func (d *ParkingSpotDLL) AddToEnd(node *ParkingSpotDLLNode) {
	newNode := node
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
}
func (d *ParkingSpotDLL) Front() *ParkingSpotDLLNode {
	return d.head
}

func (d *ParkingSpotDLL) MoveNodeToEnd(node *ParkingSpotDLLNode) {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
	if d.tail == node {
		d.tail = prev
	}
	if d.head == node {
		d.head = next
	}
	node.next = nil
	node.prev = nil
	d.len--
	d.AddToEnd(node)
}

func (d *ParkingSpotDLL) MoveNodeToFront(node *ParkingSpotDLLNode) {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
	if d.tail == node {
		d.tail = prev
	}
	if d.head == node {
		d.head = next
	}
	node.next = nil
	node.prev = nil
	d.len--
	d.AddToFront(node)
}

func (d *ParkingSpotDLL) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("Location = %v, parkingType = %s, floor = %d full =%t\n", temp.pSpot.getLocation(), temp.pSpot.getParkingSpotType().toString(), temp.pSpot.getFloor(), temp.pSpot.isFull())
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *ParkingSpotDLL) Size() int {
	return d.len
}
