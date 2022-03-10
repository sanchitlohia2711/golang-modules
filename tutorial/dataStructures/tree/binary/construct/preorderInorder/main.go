package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	lenOfTree := len(preorder)

	current := 0
	newRoot := buildTreeUtil(preorder, inorder, &current, 0, lenOfTree-1)

	return newRoot
}

func buildTreeUtil(preorder []int, inorder []int, current *int, low, high int) *TreeNode {
	if low > high {
		return nil
	}

	if low == high {
		rootNode := &TreeNode{Val: preorder[*current]}
		(*current)++
		return rootNode
	}

	rootNode := &TreeNode{Val: preorder[*current]}
	rootValue := preorder[*current]
	(*current)++

	var rootIndex int
	for i := low; i <= high; i++ {
		if inorder[i] == rootValue {
			rootIndex = i
		}
	}

	rootNode.Left = buildTreeUtil(preorder, inorder, current, low, rootIndex-1)
	rootNode.Right = buildTreeUtil(preorder, inorder, current, rootIndex+1, high)

	return rootNode
}

func main() {
	inorder := []int{4, 2, 1, 5, 3, 6}
	preorder := []int{1, 2, 4, 3, 5, 6}

	root := buildTree(preorder, inorder)
	fmt.Printf("root: %d\n", root.Val)
	fmt.Printf("root.Left: %d\n", root.Left.Val)
	fmt.Printf("root.Left.Left: %d\n", root.Left.Left.Val)
	fmt.Printf("root.Right: %d\n", root.Right.Val)
	fmt.Printf("root.Right.Left: %d\n", root.Right.Left.Val)
	fmt.Printf("root.Right.Right: %d\n", root.Right.Right.Val)
}
