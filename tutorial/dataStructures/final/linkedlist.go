package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(2)
	first.AddFront(4)
	first.AddFront(3)

	second := initList()
	second.AddFront(5)
	second.AddFront(6)
	second.AddFront(4)
}

func initList() *SingleList {
	return &SingleList{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type SingleList struct {
	Len  int
	Head *ListNode
}

func (s *SingleList) AddFront(num int) {
	ele := &ListNode{
		Val: num,
	}
	if s.Head == nil {
		s.Head = ele
	} else {
		ele.Next = s.Head
		s.Head = ele
	}
	s.Len++
}

func (s *SingleList) AddBack(num int) {
	ele := &ListNode{
		Val: num,
	}
	if s.Head == nil {
		s.Head = ele
	} else {
		current := s.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = ele
	}
	s.Len++
}

func (s *SingleList) RemoveFront() error {
	if s.Head == nil {
		return fmt.Errorf("List is empty")
	}
	s.Head = s.Head.Next
	s.Len--
	return nil
}

func (s *SingleList) RemoveBack() error {
	if s.Head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}
	var prev *ListNode
	current := s.Head
	for current.Next != nil {
		prev = current
		current = current.Next
	}
	if prev != nil {
		prev.Next = nil
	} else {
		s.Head = nil
	}
	s.Len--
	return nil
}

func (s *SingleList) Front() (int, error) {
	if s.Head == nil {
		return 0, fmt.Errorf("Single List is empty")
	}
	return s.Head.Val, nil
}

func (s *SingleList) Size() int {
	return s.Len
}
