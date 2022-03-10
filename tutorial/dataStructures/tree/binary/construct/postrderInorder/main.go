package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	lenTree := len(inorder)

	index := lenTree - 1
	return buildTreeUtil(inorder, postorder, &index, 0, lenTree-1)
}

func buildTreeUtil(inorder []int, postorder []int, index *int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	if low == high {
		currentIndexValue := postorder[*index]
		(*index)--
		return &TreeNode{Val: currentIndexValue}
	}

	currentIndexValue := postorder[*index]
	(*index)--

	root := &TreeNode{Val: currentIndexValue}

	mid := 0
	for i := low; i <= high; i++ {
		if inorder[i] == currentIndexValue {
			mid = i
		}
	}

	root.Right = buildTreeUtil(inorder, postorder, index, mid+1, high)
	root.Left = buildTreeUtil(inorder, postorder, index, low, mid-1)
	return root

}

func main() {
	inorder := []int{4, 2, 1, 5, 3, 6}
	postorder := []int{4, 2, 5, 6, 3, 1}

	root := buildTree(inorder, postorder)
	fmt.Printf("root: %d\n", root.Val)
	fmt.Printf("root.Left: %d\n", root.Left.Val)
	fmt.Printf("root.Left.Left: %d\n", root.Left.Left.Val)
	fmt.Printf("root.Right: %d\n", root.Right.Val)
	fmt.Printf("root.Right.Left: %d\n", root.Right.Left.Val)
	fmt.Printf("root.Right.Right: %d\n", root.Right.Right.Val)
}
