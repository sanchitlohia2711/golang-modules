package main

import "fmt"

type node struct {
	data string
	next *node
}

type singlyLinkedList struct {
	len  int
	head *node
}

func initList() *singlyLinkedList {
	return &singlyLinkedList{}
}

func (s *singlyLinkedList) AddFront(data string) {
	node := &node{
		data: data,
	}

	if s.head == nil {
		s.head = node
	} else {
		node.next = s.head
		s.head = node
	}
	s.len++
	return
}

func (s *singlyLinkedList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}

//Function to convert singly linked list to circular linked list
func (s *singlyLinkedList) ToCircular() {
	current := s.head
	for current.next != nil {
		current = current.next
	}
	current.next = s.head
}

func (s *singlyLinkedList) IsCircular() bool {
	if s.head == nil {
		return true
	}
	current := s.head.next
	for current.next != nil && current != s.head {
		current = current.next
	}
	return current == s.head
}

func main() {
	singleList := initList()
	fmt.Printf("AddFront: D\n")
	singleList.AddFront("D")
	fmt.Printf("AddFront: C\n")
	singleList.AddFront("C")
	fmt.Printf("AddFront: B\n")
	singleList.AddFront("B")
	fmt.Printf("AddFront: A\n")
	singleList.AddFront("A")

	err := singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	isCircular := singleList.IsCircular()
	fmt.Printf("Before: Is Circular: %t\n", isCircular)
	fmt.Printf("Size: %d\n", singleList.len)

	singleList.ToCircular()

	isCircular = singleList.IsCircular()
	fmt.Printf("After: Is Circular: %t\n", isCircular)
}
