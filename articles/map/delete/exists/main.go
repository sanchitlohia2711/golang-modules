package main

import "fmt"

func main() {
	sample := make(map[string]int)
	sample["a"] = 1
	sample["b"] = 2
	sample["c"] = 3

	fmt.Println(sample)
	delete(sample, "a")

	fmt.Println(sample)
}
