package main

import (
	"fmt"
	"math"
)

func main() {
	output := firstMissingPositive([]int{3, 2, -2})
	fmt.Println(output)

	output = firstMissingPositive([]int{-3, -2, -1})
	fmt.Println(output)

	output = firstMissingPositive([]int{7, 8, 9, 11, 12})
	fmt.Println(output)

	output = firstMissingPositive([]int{1, 2, -1})
	fmt.Println(output)

	output = firstMissingPositive([]int{1, 2, 3})
	fmt.Println(output)

	output = firstMissingPositive([]int{1, 1})
	fmt.Println(output)
}

func firstMissingPositive(nums []int) int {

	onlyPositiveNumsArray, k := segregate((nums))

	for i := 0; i < k; i++ {
		value := int(math.Abs(float64(onlyPositiveNumsArray[i])))

		if value > 0 && value <= k {
			if onlyPositiveNumsArray[value-1] > 0 {
				onlyPositiveNumsArray[value-1] = -1 * onlyPositiveNumsArray[value-1]
			}

		}
	}

	for i := 0; i < k; i++ {
		if onlyPositiveNumsArray[i] > 0 {
			return i + 1
		}
	}

	return k + 1
}

func segregate(nums []int) ([]int, int) {

	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[k] = nums[i]
			k++
		}
	}

	return nums[0:k], k

}
