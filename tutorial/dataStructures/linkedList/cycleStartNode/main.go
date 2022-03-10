package main

import "fmt"

func main() {
	first := initList()
	ele4 := first.AddFront(4)
	first.AddFront(3)
	ele2 := first.AddFront(2)
	first.AddFront(1)

	//Create cycle
	ele4.Next = ele2

	output := cycleStartNode(first.Head)
	fmt.Println(output.Val)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type SingleList struct {
	Len  int
	Head *ListNode
}

func (s *SingleList) AddFront(num int) *ListNode {
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
	return ele
}

func initList() *SingleList {
	return &SingleList{}
}
func cycleStartNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	cycleExists := false

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			cycleExists = true
			break
		}
	}

	if !cycleExists {
		return nil
	}

	cycleNode := slow

	curr := cycleNode

	lengthCycle := 1

	for curr.Next != cycleNode {
		lengthCycle++
		curr = curr.Next
	}

	curr = head

	for i := 0; i < lengthCycle; i++ {
		curr = curr.Next
	}

	for head != curr {
		head = head.Next
		curr = curr.Next
	}

	return head
}
