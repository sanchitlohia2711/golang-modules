package main

import "fmt"

func main() {
	median := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(median)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	firstArrayLen := len(nums1)
	secondArrayLen := len(nums2)

	var mid int
	i := 0
	j := 0
	var k int
	mid = (firstArrayLen+secondArrayLen)/2 + 1

	//This is the case in which the total lenght of two arrays is odd and there is only one median
	if (firstArrayLen+secondArrayLen)%2 == 1 {
		var median float64

		for k < mid {
			if i < firstArrayLen && j < secondArrayLen {
				if nums1[i] <= nums2[j] {
					median = float64(nums1[i])
					i++
					k++
				} else {
					median = float64(nums2[j])
					j++
					k++
				}
			} else if i < firstArrayLen {
				median = float64(nums1[i])
				i++
				k++
			} else {
				median = float64(nums2[j])
				j++
				k++
			}

		}
		return median
	} else { //This is the case in which the total lenght of two arrays is even and there is only two medians. We need to return average of these two medians
		var median1 float64
		var median2 float64

		for k < mid {
			median1 = median2
			if i < firstArrayLen && j < secondArrayLen {
				if nums1[i] <= nums2[j] {
					median2 = float64(nums1[i])
					i++
					k++
				} else {
					median2 = float64(nums2[j])
					j++
					k++
				}
			} else if i < firstArrayLen {
				median2 = float64(nums1[i])
				i++
				k++
			} else {
				median2 = float64(nums2[j])
				j++
				k++
			}

		}
		return (median1 + median2) / 2
	}
}
