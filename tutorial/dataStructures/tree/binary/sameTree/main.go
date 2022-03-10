package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	root1 := TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 3}
	root1.Right.Left = &TreeNode{Val: 5}
	root1.Right.Right = &TreeNode{Val: 6}

	root2 := TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Right = &TreeNode{Val: 3}
	root2.Right.Left = &TreeNode{Val: 5}
	root2.Right.Right = &TreeNode{Val: 6}

	output := isSameTree(&root1, &root2)
	fmt.Println(output)

	root1 = TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}

	root2 = TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 3}

	output = isSameTree(&root1, &root2)
	fmt.Println(output)

}
