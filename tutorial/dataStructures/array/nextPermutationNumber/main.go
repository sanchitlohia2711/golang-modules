package main

import (
	"fmt"
	"sort"
)

func main() {
	nextPermutation([]int{1, 2, 3})

	nextPermutation([]int{2, 1})

	nextPermutation([]int{1, 3, 5, 4, 1})

	nextPermutation([]int{1, 3, 2})
}

func nextPermutation(nums []int) {

	numsLen := len(nums)
	first := -1
	second := -1

	for i, j := numsLen-2, numsLen-1; i >= 0; {
		if nums[i] < nums[j] {
			first = i
			second = j
			break
		} else {
			i--
			j--
		}
	}

	if !(first == -1) {
		smallestGreaterIndex := second
		for i := second + 1; i < numsLen; i++ {
			if nums[i] > nums[first] && nums[i] < nums[smallestGreaterIndex] {
				smallestGreaterIndex = i
			}
		}
		nums[first], nums[smallestGreaterIndex] = nums[smallestGreaterIndex], nums[first]

		sort.Slice(nums[second:numsLen], func(i, j int) bool {
			return nums[second+i] < nums[second+j]
		})
	}

	fmt.Println(nums)

}
