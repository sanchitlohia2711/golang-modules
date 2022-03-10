package main

import "fmt"

type node struct {
	data string
	prev *node
	next *node
}

type doublyLinkedList struct {
	len  int
	tail *node
	head *node
}

type cache struct {
	dll          *doublyLinkedList
	storage      map[string]*node
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}
