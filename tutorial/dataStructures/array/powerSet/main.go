package main

import (
	"fmt"
	"math"
)

func subsets(nums []int) [][]int {

	lengthNums := len(nums)
	powerSetLength := int(math.Pow(2, float64(lengthNums)))
	output := make([][]int, 0)

	for i := 0; i < powerSetLength; i++ {

		result := make([]int, 0)
		for j := 0; j < lengthNums; j++ {
			val := int(i) & int(1<<j)
			if val != 0 {
				result = append(result, nums[j])
			}
		}

		output = append(output, result)
	}

	return output
}

func main() {
	output := subsets([]int{1, 2})
	fmt.Println(output)

	output = subsets([]int{1, 2, 3})
	fmt.Println(output)
}
