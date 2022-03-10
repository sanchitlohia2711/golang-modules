package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(4)
	first.AddFront(3)
	first.AddFront(2)
	first.AddFront(1)

	first.Head.Traverse()
	temp := ReverseKGroup(first.Head, 3)
	fmt.Println()
	temp.Traverse()

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

func (s *SingleList) Reverse() {

	curr := s.Head
	var prev *ListNode
	var next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	s.Head = prev
}

func ReverseKGroup(head *ListNode, k int) *ListNode {

	curr := head
	var prev *ListNode
	var next *ListNode

	i := 0

	for curr != nil && i < k {
		i++
		curr = curr.Next
	}
	if i == k {
		curr = head
	} else {
		return head
	}

	i = 0
	for curr != nil && i < k {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		i++
	}

	head.Next = ReverseKGroup(curr, k)
	return prev
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
