package main

import (
	"fmt"
)

func removeElement(nums []int, val int) []int {
	lenNums := len(nums)

	k := 0

	for i := 0; i < lenNums; {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
		i++
	}
	return nums[0:k]
}

func main() {
	output := removeElement([]int{1, 4, 2, 5, 4}, 4)
	fmt.Println(output)

	output = removeElement([]int{1, 2, 3}, 3)
	fmt.Println(output)
}
