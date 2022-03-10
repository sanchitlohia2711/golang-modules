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
	removeNthFromEnd(first.Head, 2)
	fmt.Println("")
	first.Head.Traverse()

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := lengthOfList(head)

	numberFront := size - n + 1

	if numberFront < 1 {
		return head
	}

	if numberFront == 1 {
		return head.Next
	}

	if size == 1 {
		return nil
	}

	//Get to the numberFront-1 node
	curr := head
	for i := 0; i < numberFront-2; i++ {
		curr = curr.Next
	}

	prev := curr

	nodeToDelete := curr.Next

	if nodeToDelete != nil {
		nodeToDeleteNext := nodeToDelete.Next
		prev.Next = nodeToDeleteNext
	} else {
		prev.Next = nil
	}

	return head
}

func lengthOfList(head *ListNode) int {
	size := 0
	for head != nil {
		size = size + 1
		head = head.Next
	}
	return size
}
