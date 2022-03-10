package main

import "fmt"

type LRUCache struct {
	doublyLinkedList *doublyLinkedList
	storage          map[string]*node
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

func (this *LRUCache) Put(key, value string) {
	if this.capacity == this.maxCapacity {
		evictedKey := this.evict()
		delete(this.storage, evictedKey)
	}
	node := &node{key: key, value: value}
	this.storage[key] = node
	this.evictionAlgo.set(node, this)
	this.capacity++
}

func (this *LRUCache) Get(key string) string {
	node_ptr, ok := this.storage[key]
	if ok {
		this.evictionAlgo.get(node_ptr, this)
		return (*node_ptr).value
	}
	return ""
}

func (this *LRUCache) evict() string {
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
