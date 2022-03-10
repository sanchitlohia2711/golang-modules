package main

import "fmt"

func main() {
	sortNums([]int{2, 0, 2, 1, 1, 0})
}

func sortNums(nums []int) {
	zeroCount := 0
	oneCount := 0
	twoCount := 0

	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			zeroCount++
		case 1:
			oneCount++
		case 2:
			twoCount++
		}
	}
	counter := 0
	for i := 0; i < zeroCount; i++ {
		nums[counter] = 0
		counter++
	}
	for i := 0; i < oneCount; i++ {
		nums[counter] = 1
		counter++
	}
	for i := 0; i < twoCount; i++ {
		nums[counter] = 2
		counter++
	}
	fmt.Println(nums)
}
