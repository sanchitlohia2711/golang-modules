package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{3, 4, 2, 1}
	start := 2
	end := 4
	sort.Slice(nums[start:end], func(i, j int) bool {
		return nums[start+i] < nums[start+j]
	})
	fmt.Println(nums)
}
