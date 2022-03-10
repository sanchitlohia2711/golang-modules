package main

import "fmt"

func main() {
	output := searchRange([]int{1, 2, 2, 5}, 2)
	fmt.Println(output)

	output = searchRange([]int{1, 2, 5}, 2)
	fmt.Println(output)

	output = searchRange([]int{}, 1)
	fmt.Println(output)

	// output = searchRange([]int{5, 7, 7, 8, 8, 10}, 7)
	// fmt.Println(output)

	// output = search([]int{4, 5, 6, 7, 0, 1, 2}, 2)
	// fmt.Println(output)

	// output = search([]int{4, 5, 6, 7, 0, 1, 2}, 4)
	// fmt.Println(output)

	// output = search([]int{4, 5, 6, 7, 0, 1, 2}, 5)
	// fmt.Println(output)

	// output = search([]int{4, 5, 6, 7, 0, 1, 2}, 6)
	// fmt.Println(output)

	// output = search([]int{4, 5, 6, 7, 0, 1, 2}, 7)
	// fmt.Println(output)

	// output = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	// fmt.Println(output)

	// output = search([]int{1, 2}, 3)
	// fmt.Println(output)
}

func searchRange(nums []int, target int) []int {

	output := make([]int, 2)
	output[0] = findLeftPivot(nums, 0, len(nums)-1, target, len(nums))

	output[1] = findRightPivot(nums, 0, len(nums)-1, target, len(nums))

	return output
}

func findLeftPivot(nums []int, start, end, target, len int) int {
	if start > end {
		return -1
	}

	if start == end && nums[start] == target {
		return start
	}

	mid := (start + end) / 2

	if (mid == 0 || nums[mid-1] < nums[mid]) && nums[mid] == target {
		return mid
	}

	if target <= nums[mid] {
		return findLeftPivot(nums, start, mid-1, target, len)
	}

	return findLeftPivot(nums, mid+1, end, target, len)

}

func findRightPivot(nums []int, start, end, target, len int) int {
	if start > end {
		return -1
	}

	if start == end && nums[start] == target {
		return start

	}

	mid := (start + end) / 2

	if mid+1 <= end && nums[mid] == target && nums[mid] < nums[mid+1] {
		return mid
	}

	if (mid == len-1 || nums[mid] < nums[mid+1]) && nums[mid] == target {
		return mid - 1
	}

	if target >= nums[mid] {
		return findRightPivot(nums, mid+1, end, target, len)
	}

	return findRightPivot(nums, start, mid-1, target, len)

}
