package main

import "fmt"

func main() {
	median := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(median)
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	firstArrayLen := len(nums1)
	secondArrayLen := len(nums2)

	var mid int

	if (firstArrayLen+secondArrayLen)%2 == 0 {
		mid = (firstArrayLen + secondArrayLen) / 2
	} else {
		mid = (firstArrayLen+secondArrayLen)/2 + 1
	}

	if mid == 0 {
		if firstArrayLen == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}

	var firstMedian float64
	var secondMedian float64
	i := 0
	j := 0
	k := 0
	for i < firstArrayLen && j < secondArrayLen {
		if nums1[i] <= nums2[j] {
			k++
			if k == mid {
				firstMedian = float64(nums1[i])
			} else if k == mid+1 {
				secondMedian = float64(nums1[i])
				break
			}
			i++
		} else if nums1[i] > nums2[j] {
			k++
			if k == mid {
				firstMedian = float64(nums2[j])
			} else if k == mid+1 {
				secondMedian = float64(nums2[j])
				break
			}
			j++
		}
	}

	for i < firstArrayLen {
		if k > mid+1 {
			break
		}

		k++
		if k == mid {
			firstMedian = float64(nums1[i])
		} else if k == mid+1 {
			secondMedian = float64(nums1[i])
			break
		}
		i++
	}

	for j < secondArrayLen {
		if k > mid+1 {
			break
		}
		k++
		if k == mid {
			firstMedian = float64(nums2[j])
		} else if k == mid+1 {
			secondMedian = float64(nums2[j])
			break
		}
		j++
	}

	if (firstArrayLen+secondArrayLen)%2 == 0 {
		return (firstMedian + secondMedian) / 2
	} else {
		return firstMedian
	}

}
