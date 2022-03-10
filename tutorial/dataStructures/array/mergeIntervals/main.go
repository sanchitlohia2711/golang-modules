package main

import (
	"fmt"
	"sort"
)

func main() {
	output := merge([][]int{{1, 4}, {8, 10}, {9, 12}, {3, 5}})
	fmt.Println(output)

	output = merge([][]int{{1, 4}, {4, 5}})
	fmt.Println(output)

	output = merge([][]int{{2, 2}, {2, 2}})
	fmt.Println(output)

	output = merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}})
	fmt.Println(output)
}

type intervalsArray [][]int

func (intA intervalsArray) Len() int {
	return len(intA)
}

func (intA intervalsArray) Swap(i, j int) {
	intA[i], intA[j] = intA[j], intA[i]
}

func (intA intervalsArray) Less(i, j int) bool {
	return intA[i][0] < intA[j][0]
}

func merge(intervals [][]int) [][]int {

	intA := intervalsArray(intervals)

	sort.Sort(intA)

	intervalsSorted := [][]int(intA)
	//fmt.Println(intervalsSorted)

	var output [][]int
	currentIntervalStart := intervalsSorted[0][0]
	currentIntervalEnd := intervalsSorted[0][1]
	for j := 1; j < len(intervalsSorted); j++ {
		if currentIntervalEnd >= intervalsSorted[j][0] {
			if intervalsSorted[j][1] > currentIntervalEnd {
				currentIntervalEnd = intervalsSorted[j][1]
			}
		} else {
			output = append(output, []int{currentIntervalStart, currentIntervalEnd})
			currentIntervalStart = intervalsSorted[j][0]
			currentIntervalEnd = intervalsSorted[j][1]
		}
	}
	output = append(output, []int{currentIntervalStart, currentIntervalEnd})
	return output

}
