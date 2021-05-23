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

func (s *singlyLinkedList) Removekth(k int) error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	if k > s.len {
		return fmt.Errorf("Err: Given number is greater than linked list length")
	}
	if k == 1 {
		fmt.Println("Node to be removed is front node")
		s.head = s.head.next
	} else {
		var prev *node
		current := s.head
		for i := 1; i < k; i++ {
			prev = current
			current = current.next
		}
		prev.next = current.next
	}
	s.len--
	return nil
}

func main() {
	// Initiaise singly linked list
	singleList := initList()
	// Adding nodes in linked list from front
	fmt.Printf("AddFront: F\n")
	singleList.AddFront("F")
	fmt.Printf("AddFront: E\n")
	singleList.AddFront("E")
	fmt.Printf("AddFront: D\n")
	singleList.AddFront("D")
	fmt.Printf("AddFront: C\n")
	singleList.AddFront("C")
	fmt.Printf("AddFront: B\n")
	singleList.AddFront("B")
	fmt.Printf("AddFront: A\n")
	singleList.AddFront("A")

	fmt.Println("Traversal")
	err := singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("\nRemoving 2nd node from the front which is B")
	err = singleList.Removekth(2)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Traversal after 2nd node is removed from the front")
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("\nRemoving 5th node from the front which is F now")
	err = singleList.Removekth(5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Traversal after 5th node is removed from the front")
	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Trying to delete node from a place greater than list size
	fmt.Println("\nTrying to delete node from a place greater than list size")
	err = singleList.Removekth(5)
	if err != nil {
		fmt.Println(err.Error())
	}

}
