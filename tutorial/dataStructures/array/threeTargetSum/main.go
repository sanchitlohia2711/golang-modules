package main

import (
	"fmt"
)

func main() {
	output := threeSumTarget([]int{1, 2, 3, 4, 1}, 6)
	fmt.Println(output)
}

func threeSumTarget(nums []int, target int) [][]int {

	numsLength := len(nums)
	var results [][]int
	for k := 0; k < numsLength-2; k++ {

		for i := k + 1; i < numsLength; i++ {
			if numsMap[target-(nums[i]+nums[k])] > 0 {
				result := []int{nums[k], target - (nums[i] + nums[k]), nums[i]}
				results = append(results, result)
			}
			numsMap[nums[i]] = i
		}
	}

	return results
	//return removeDuplicates(results)
}

// func removeDuplicates(inputs [][]int) [][]int {
// 	numberMap := make(map[string]int)
// 	var results [][]int
// 	for _, input := range inputs {
// 		sort.Slice(input, func(i, j int) bool {
// 			return input[i] < input[j]
// 		})
// 		var numString string
// 		for i := range input {
// 			numString = numString + strconv.Itoa(input[i])
// 		}
// 		_, ok := numberMap[numString]
// 		if !ok {
// 			results = append(results, input)
// 		}
// 		numberMap[numString] = 1
// 	}
// 	return results
// }
