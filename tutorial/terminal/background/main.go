package main

import "fmt"

func main() {
	output := twoTargetSums([]int{2, 5, 1, 3}, 4)
	fmt.Println(output)
}

func twoTargetSums(nums []int, target int) []int {
	numberMap := make(map[int]int)
	output := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		val, ok := numberMap[target-nums[i]]
		if ok {
			output[0] = val
			output[1] = i
			return output
		} else {
			numberMap[nums[i]] = i
		}
	}
	return output
}
