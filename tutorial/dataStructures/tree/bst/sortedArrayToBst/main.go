package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	return sortedArrayToBSTUtil(nums, 0, length-1)
}

func sortedArrayToBSTUtil(nums []int, first int, last int) *TreeNode {

	if first > last {
		return nil
	}

	if first == last {
		return &TreeNode{
			Val: nums[first],
		}
	}

	mid := (first + last) / 2

	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = sortedArrayToBSTUtil(nums, first, mid-1)
	root.Right = sortedArrayToBSTUtil(nums, mid+1, last)
	return root
}

func traverseInorder(root *TreeNode) {
	if root == nil {
		return
	}

	traverseInorder(root.Left)
	fmt.Println(root.Val)
	traverseInorder(root.Right)
}

func main() {
	root := sortedArrayToBST([]int{1, 2, 3, 4, 5})
	traverseInorder(root)

}
