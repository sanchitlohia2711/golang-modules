package main

import "fmt"

const (
	UintSize = 32 << (^uint(0) >> 32 & 1)
	MaxInt   = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt   = -MaxInt - 1
)

func main() {
	input := []int{4, 5, -3}
	output := maxSubArray(input)
	fmt.Println(output)

	input = []int{1, 2, -4, 4, 1}
	output = maxSubArray(input)
	fmt.Println(output)
}

func maxSubArray(nums []int) int {
	lenNums := len(nums)

	currentMax := 0
	max := MinInt
	for i := 0; i < lenNums; i++ {
		currentMax = currentMax + nums[i]
		if currentMax > max {
			max = currentMax
		}

		if currentMax < 0 {
			currentMax = 0
		}

	}
	return max
}
