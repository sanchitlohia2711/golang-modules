package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(5)
	first.AddFront(4)
	first.AddFront(3)
	first.AddFront(2)
	first.AddFront(1)

	first.Head.Traverse()

	head := rotateRight(first.Head, 2)
	fmt.Println("")
	head.Traverse()
	fmt.Println("")

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
		fmt.Print(l.Val)
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

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}

	if k == 0 {
		return head
	}

	lenList := 0

	curr := head
	var last *ListNode

	for curr != nil {
		lenList++
		last = curr
		curr = curr.Next
	}

	k = k % lenList
	if k == 0 {
		return head
	}

	curr = head

	newHeadIndex := lenList - k

	var prev *ListNode

	for i := 0; i < newHeadIndex; i++ {
		prev = curr
		curr = curr.Next
	}

	newHeadNode := prev.Next

	prev.Next = nil

	last.Next = head
	return newHeadNode
}
