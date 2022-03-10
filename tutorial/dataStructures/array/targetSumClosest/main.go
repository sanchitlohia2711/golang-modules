package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	output := threeSumClosest([]int{0, 2, 3, -1}, 6)
	fmt.Println(output)
}

func threeSumClosest(nums []int, target int) int {

	numsLength := len(nums)

	closestSum := 0
	nearest := 10000

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for k := 0; k < numsLength-2; k++ {
		i := k + 1
		j := numsLength - 1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			absSum := int(math.Abs(float64(target - sum)))

			if absSum < nearest {
				nearest = absSum
				closestSum = sum
			}

			if nums[k]+nums[i]+nums[j] > target {
				j--
			} else {
				i++
			}
		}
	}

	return closestSum
}
