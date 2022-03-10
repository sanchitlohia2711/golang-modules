package main

import "fmt"

func canJump(nums []int) bool {
	lenNums := len(nums)
	canJumpB := make([]bool, lenNums)

	canJumpB[0] = true

	for i := 0; i < lenNums; i++ {

		if canJumpB[i] {
			valAtCurrIndex := nums[i]
			for k := 1; k <= valAtCurrIndex && i+k < lenNums; k++ {
				canJumpB[i+k] = true
			}
		}

	}

	return canJumpB[lenNums-1]
}

func main() {
	input := []int{2, 3, 1, 1, 4}

	canJumpOrNot := canJump(input)
	fmt.Println(canJumpOrNot)

	input = []int{3, 2, 1, 0, 4}

	canJumpOrNot = canJump(input)
	fmt.Println(canJumpOrNot)

}
