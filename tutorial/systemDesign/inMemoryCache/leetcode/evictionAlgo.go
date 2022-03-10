package main

type evictionAlgo interface {
	evict(c *LRUCache) int
	get(node *node, c *LRUCache)
	set(node *node, c *LRUCache)
}

func createEvictioAlgo(algoType string) evictionAlgo {
	if algoType == "fifo" {
		return &fifo{}
	} else if algoType == "lru" {
		return &lru{}
	}

	return nil
}
