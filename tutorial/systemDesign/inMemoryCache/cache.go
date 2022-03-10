package main

import "fmt"

type Cache struct {
	doublyLinkedList *doublyLinkedList
	storage          map[string]*node
	evictionAlgo     evictionAlgo
	capacity         int
	maxCapacity      int
}

func initCache(evictionAlgo evictionAlgo, maxCapacity int) Cache {
	storage := make(map[string]*node)
	return Cache{
		doublyLinkedList: &doublyLinkedList{},
		storage:          storage,
		evictionAlgo:     evictionAlgo,
		capacity:         0,
		maxCapacity:      maxCapacity,
	}
}

func (this *Cache) setEvictionAlgo(e evictionAlgo) {
	this.evictionAlgo = e
}

func (this *Cache) set(key, value string) {
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

func (this *Cache) get(key string) string {
	node_ptr, ok := this.storage[key]
	if ok {
		this.evictionAlgo.get(node_ptr, this)
		return (*node_ptr).value
	}
	return ""
}

func (this *Cache) evict() string {
	key := this.evictionAlgo.evict(this)
	this.capacity--
	return key
}

func (this *Cache) print() {
	for k, v := range this.storage {
		fmt.Printf("key :%s value: %s\n", k, (*v).value)
	}
	this.doublyLinkedList.TraverseForward()
}
