package main

import "fmt"

type lru struct {
}

func (l *lru) evict(c *LRUCache) int {
	key := c.doublyLinkedList.Front().key
	fmt.Println("Evicting by lru strtegy. Evicted Node Key: %s: ", key)
	c.doublyLinkedList.RemoveFromFront()
	return key
}

func (l *lru) get(node *node, c *LRUCache) {
	fmt.Println("Shuffling doubly linked list due to get operation")
	c.doublyLinkedList.MoveNodeToEnd(node)
}

func (l *lru) set(node *node, c *LRUCache) {
	fmt.Println("Shuffling doubly linked list due to set operation")
	c.doublyLinkedList.AddToEnd(node)
}
