package main

import "fmt"

func main() {
	input := []int{1, 1, 1, 2}
	output := removeDuplicates(input)
	fmt.Println(output)

	input = []int{1, 2, 3, 3}
	output = removeDuplicates(input)
	fmt.Println(output)
}

func removeDuplicates(nums []int) []int {
	lenArray := len(nums)

	k := 0
	for i := 0; i < lenArray; {
		nums[k] = nums[i]
		k++
		for i+1 < lenArray && nums[i] == nums[i+1] {
			i++
		}
		i++
	}

	return nums[0:k]
}
