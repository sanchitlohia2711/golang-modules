package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	lengthCandidates := len(candidates)
	current_sum_array := make([]int, 0)
	output := make([][]int, 0)
	combinationSumUtil(candidates, lengthCandidates, 0, 0, 0, target, current_sum_array, &output)
	return output
}

func combinationSumUtil(candidates []int, lengthCandidates, index, current_sum_index, current_sum, target int, current_sum_array []int, output *[][]int) {

	if index >= lengthCandidates {
		return
	}

	if current_sum > target {
		return
	}

	if current_sum == target {
		var o []int
		for i := 0; i < current_sum_index; i++ {
			o = append(o, current_sum_array[i])
		}
		*output = append(*output, o)
		return
	}

	//Exclude
	combinationSumUtil(candidates, lengthCandidates, index+1, current_sum_index, current_sum, target, current_sum_array, output)

	//Include
	current_sum_array = append(current_sum_array, candidates[index])

	combinationSumUtil(candidates, lengthCandidates, index, current_sum_index+1, current_sum+candidates[index], target, current_sum_array, output)
}

func main() {
	output := combinationSum([]int{3, 4, 10, 11}, 10)
	fmt.Println(output)
}
