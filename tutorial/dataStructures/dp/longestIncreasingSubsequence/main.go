package main

import "fmt"

func lengthOfLIS(nums []int) int {

	lenNums := len(nums)
	lis := make([]int, lenNums)

	for i := 0; i < lenNums; i++ {
		lis[i] = 1
	}

	for i := 1; i < lenNums; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && lis[i] < (lis[j]+1) {
				lis[i] = lis[j] + 1
			}
		}
	}

	max := 0

	for i := 0; i < lenNums; i++ {
		if lis[i] > max {
			max = lis[i]
		}
	}

	return max
}

func main() {
	output := lengthOfLIS([]int{1, 5, 7, 6})
	fmt.Println(output)

	output = lengthOfLIS([]int{3, 2, 1})
	fmt.Println(output)
}
