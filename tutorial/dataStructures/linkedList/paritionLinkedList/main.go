package main

import "fmt"

func main() {
	first := initList()
	first.AddFront(1)
	first.AddFront(3)
	first.AddFront(5)
	first.AddFront(4)

	first.Head.Traverse()
	newHead := partition(first.Head, 2)
	fmt.Println("")
	newHead.Traverse()

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

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	curr := head

	var prev *ListNode

	for curr != nil {
		if curr.Val >= x {
			break
		}
		prev = curr
		curr = curr.Next
	}

	if curr == nil {
		return head
	}

	firstLargeValueNode := curr

	prev2 := firstLargeValueNode
	for curr != nil {
		if curr.Val < x {
			prev2.Next = curr.Next
			if prev != nil {
				prev.Next = curr
				prev = prev.Next
				prev.Next = firstLargeValueNode
			} else {
				if head == firstLargeValueNode {
					head = curr
				}
				curr.Next = firstLargeValueNode
				prev = curr
			}
		}
		prev2 = curr
		curr = curr.Next
	}

	return head
}
