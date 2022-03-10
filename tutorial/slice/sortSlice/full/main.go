package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{3, 4, 2, 1}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
}
