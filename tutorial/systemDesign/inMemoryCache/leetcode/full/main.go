package main

import "fmt"

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type doublyLinkedList struct {
	len  int
	tail *node
	head *node
}

func initDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func (d *doublyLinkedList) AddToFront(key, value int) {
	newNode := &node{
		key:   key,
		value: value,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *doublyLinkedList) RemoveFromFront() {
	if d.head == nil {
		return
	} else if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
	}
	d.len--
}

func (d *doublyLinkedList) AddToEnd(node *node) {
	newNode := node
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
}
func (d *doublyLinkedList) Front() *node {
	return d.head
}

func (d *doublyLinkedList) MoveNodeToEnd(node *node) {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
	if d.tail == node {
		d.tail = prev
	}
	if d.head == node {
		d.head = next
	}
	node.next = nil
	node.prev = nil
	d.len--
	d.AddToEnd(node)
}

func (d *doublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		fmt.Printf("key = %v, value = %v, prev = %v, next = %v\n", temp.key, temp.value, temp.prev, temp.next)
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *doublyLinkedList) Size() int {
	return d.len
}

type evictionAlgo interface {
	evict(c *LRUCache) int
	get(node *node, c *LRUCache)
	set(node *node, c *LRUCache)
	set_overwrite(node *node, value int, c *LRUCache)
}

func createEvictioAlgo(algoType string) evictionAlgo {
	if algoType == "fifo" {
		return &fifo{}
	} else if algoType == "lru" {
		return &lru{}
	}

	return nil
}

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

func (l *fifo) set_overwrite(node *node, value int, c *LRUCache) {
	fmt.Println("Shuffling doubly linked list due to set_overwrite operation")
}

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

func (l *lru) set_overwrite(node *node, value int, c *LRUCache) {
	fmt.Println("Shuffling doubly linked list due to set_overwrite operation")
	node.value = value
	c.doublyLinkedList.MoveNodeToEnd(node)
}

type LRUCache struct {
	doublyLinkedList *doublyLinkedList
	storage          map[int]*node
	evictionAlgo     evictionAlgo
	capacity         int
	maxCapacity      int
}

func Constructor(capacity int) LRUCache {
	storage := make(map[int]*node)
	lru := createEvictioAlgo("lru")
	return LRUCache{
		doublyLinkedList: &doublyLinkedList{},
		storage:          storage,
		evictionAlgo:     lru,
		capacity:         0,
		maxCapacity:      capacity,
	}
}

func (this *LRUCache) setEvictionAlgo(e evictionAlgo) {
	this.evictionAlgo = e
}

func (this *LRUCache) Put(key int, value int) {
	node_ptr, ok := this.storage[key]
	if ok {
		this.evictionAlgo.set_overwrite(node_ptr, value, this)
		return
	}
	if this.capacity == this.maxCapacity {
		evictedKey := this.evict()
		delete(this.storage, evictedKey)
	}
	node := &node{key: key, value: value}
	this.storage[key] = node
	this.evictionAlgo.set(node, this)
	this.capacity++
}

func (this *LRUCache) Get(key int) int {
	node_ptr, ok := this.storage[key]
	if ok {
		this.evictionAlgo.get(node_ptr, this)
		return (*node_ptr).value
	}
	return -1
}

func (this *LRUCache) evict() int {
	key := this.evictionAlgo.evict(this)
	this.capacity--
	return key
}

func (this *LRUCache) print() {
	for k, v := range this.storage {
		fmt.Printf("key :%s value: %s\n", k, (*v).value)
	}
	this.doublyLinkedList.TraverseForward()
}

func main() {

}
