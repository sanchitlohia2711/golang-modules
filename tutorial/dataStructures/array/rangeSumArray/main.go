package main

import "fmt"

type NumArray struct {
	sum_nums []int
}

func initNumArray(nums []int) NumArray {
	length := len(nums)
	sum_nums := make([]int, length)

	sum_nums[0] = nums[0]

	for i := 1; i < length; i++ {
		sum_nums[i] = nums[i] + sum_nums[i-1]
	}

	return NumArray{
		sum_nums: sum_nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	leftSum := 0
	if left > 0 {
		leftSum = this.sum_nums[left-1]
	}
	return this.sum_nums[right] - leftSum
}

func main() {
	nums := []int{1, 3, 5, 6, 2}
	na := initNumArray(nums)

	output := na.SumRange(2, 4)
	fmt.Println(output)

}
