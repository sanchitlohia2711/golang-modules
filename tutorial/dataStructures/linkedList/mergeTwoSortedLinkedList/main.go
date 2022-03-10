package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(4)
	first.AddFront(3)
	first.AddFront(2)

	second := initList()
	second.AddFront(6)
	second.AddFront(5)
	second.AddFront(4)

	output := mergeTwoLists(first.Head, second.Head)
	output.Traverse()

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type SingleList struct {
	Len  int
	Head *ListNode
}

func initList() *SingleList {
	return &SingleList{}
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var output *ListNode
	var prev *ListNode

	for l1 != nil && l2 != nil {
		val := 0
		if l1.Val < l2.Val {
			val = l1.Val
			l1 = l1.Next
		} else {
			val = l2.Val
			l2 = l2.Next
		}

		if prev == nil {
			output = &ListNode{Val: val}
			prev = output
		} else {
			temp := &ListNode{Val: val}
			prev.Next = temp
			prev = temp
		}

	}

	for l1 != nil {
		if prev == nil {
			output = &ListNode{Val: l1.Val}
			prev = output
		} else {
			temp := &ListNode{Val: l1.Val}
			prev.Next = temp
			prev = temp
		}
		l1 = l1.Next
	}

	for l2 != nil {
		if prev == nil {
			output = &ListNode{Val: l2.Val}
			prev = output
		} else {
			temp := &ListNode{Val: l2.Val}
			prev.Next = temp
			prev = temp
		}
		l2 = l2.Next
	}

	return output
}
