package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(4)
	first.AddFront(3)
	first.AddFront(2)
	first.AddFront(1)

	output := swapPairs(first.Head)
	output.Traverse()
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

func swapPairs(head *ListNode) *ListNode {
	var current *ListNode
	var next *ListNode
	var prev *ListNode
	current = head

	if current != nil && current.Next != nil {
		head = current.Next
	}
	for current != nil && current.Next != nil {
		next = current.Next.Next
		if prev == nil {
			prev = current
		} else {
			prev.Next = current.Next
			prev = current
		}
		current.Next.Next = current
		current.Next = next
		current = next

	}
	return head
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
