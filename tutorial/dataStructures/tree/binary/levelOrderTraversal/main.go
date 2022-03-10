package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levelMapNode := make(map[int][]int)
	visit(root, 0, &levelMapNode)

	output := make([][]int, 0)

	i := 0
	for true {
		_, ok := levelMapNode[i]

		if ok {
			output = append(output, levelMapNode[i])
		} else {
			break
		}
		i = i + 1

	}
	return output

}

func visit(root *TreeNode, level int, levelMapNode *map[int][]int) {
	if root == nil {
		return
	}
	_, ok := (*levelMapNode)[level]
	if ok {
		(*levelMapNode)[level] = append((*levelMapNode)[level], root.Val)
	} else {
		(*levelMapNode)[level] = []int{root.Val}
	}

	if root.Left != nil {
		visit(root.Left, level+1, levelMapNode)
	}

	if root.Right != nil {
		visit(root.Right, level+1, levelMapNode)
	}
	return
}

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	output := levelOrder(&root)
	fmt.Println(output)

}
