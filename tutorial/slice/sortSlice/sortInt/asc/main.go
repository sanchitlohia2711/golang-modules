package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{3, 4, 2, 1}
	fmt.Printf("Before: %v", nums)
	sort.Ints(nums)
	fmt.Printf("\nAfter: %v", nums)
}
