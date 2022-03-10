package main

import "fmt"

func main() {
	// lru := createEvictioAlgo("lru")
	// cache := initCache(lru)
	// cache.add("a", "1")
	// cache.print()

	// cache.add("b", "2")
	// cache.print()

	// cache.add("c", "3")
	// cache.print()

	// fifo := createEvictioAlgo("fifo")
	// cache.setEvictionAlgo(fifo)

	// cache.add("d", "4")
	// cache.print()

	// cache.add("e", "5")
	// cache.print()

	lru := createEvictioAlgo("lru")
	cache := initCache(lru)
	cache.set("a", "1")
	cache.print()

	cache.set("b", "2")
	cache.print()

	cache.set("c", "3")
	cache.print()

	value := cache.get("a")
	fmt.Printf("key: a, value: %s\n", value)
	cache.print()

	cache.set("d", "4")
	cache.print()

	cache.set("e", "5")
	cache.print()
}
