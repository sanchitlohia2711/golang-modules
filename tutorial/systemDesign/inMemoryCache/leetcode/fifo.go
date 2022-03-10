package main

import "fmt"

type fifo struct {
}

func (l *fifo) evict(c *LRUCache) int {
	fmt.Println("Evicting by fifo strtegy")
	key := c.doublyLinkedList.Front().key
	c.doublyLinkedList.RemoveFromFront()
	return key
}

func (l *fifo) get(node *node, c *LRUCache) {
	fmt.Println("Shuffling doubly linked list due to get operation")
}

func (l *fifo) set(node *node, c *LRUCache) {
	fmt.Println("Shuffling doubly linked list due to set operation")
}
