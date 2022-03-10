package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	isValid, _, _ := isValidBSTUtil(root)
	return isValid
}

func isValidBSTUtil(node *TreeNode) (bool, int, int) {

	if node.Left == nil && node.Right == nil {
		return true, node.Val, node.Val
	}

	min := node.Val
	max := node.Val

	isValidLeft := true
	var leftMin, leftMax int
	if node.Left != nil {
		isValidLeft, leftMin, leftMax = isValidBSTUtil(node.Left)

		if !isValidLeft {
			return false, 0, 0
		}
		if node.Val <= leftMax {
			return false, 0, 0
		}

		min = leftMin
	}

	isValidRight := true
	var rightMin, rightMax int

	if node.Right != nil {
		isValidRight, rightMin, rightMax = isValidBSTUtil(node.Right)

		if !isValidRight {
			return false, 0, 0
		}

		if node.Val >= rightMin {
			return false, 0, 0
		}
		max = rightMax
	}

	return true, min, max
}

func minOfFour(a, b, c, d int) int {
	return int(math.Min(float64(a), math.Min(float64(b), math.Min(float64(c), float64(d)))))
}

func maxOfFour(a, b, c, d int) int {
	return int(math.Max(float64(a), math.Max(float64(b), math.Max(float64(c), float64(d)))))
}

func main() {
	root := TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	isValidBST := isValidBST(&root)
	fmt.Println(isValidBST)

}
