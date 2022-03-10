package main

import "fmt"

func jump(nums []int) int {

	lenJump := len(nums)
	minJumps := make([]int, lenJump)
	for i := 0; i < lenJump; i++ {
		minJumps[i] = -1
	}

	minJumps[0] = 0

	for i := 0; i < lenJump; i++ {
		currVal := nums[i]

		for j := 1; j <= currVal && i+j < lenJump; j++ {
			if minJumps[i+j] == -1 || minJumps[i+j] > minJumps[i]+1 {
				minJumps[i+j] = minJumps[i] + 1
			}
		}
	}

	return minJumps[lenJump-1]

}

func main() {
	input := []int{2, 3, 1, 1, 4}

	minJump := jump(input)
	fmt.Println(minJump)

	input = []int{3, 2, 1, 0, 4}

	minJump = jump(input)
	fmt.Println(minJump)

}
