package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Just keeping track of the max depth in an int pointer, and a local curDepth var.
func next(root *TreeNode, totalDepth *int, curDepth int) {
	if *totalDepth < curDepth {
		*totalDepth = curDepth
	}

	if root.Left != nil {
		next(root.Left, totalDepth, curDepth+1)
	}

	if root.Right != nil {
		next(root.Right, totalDepth, curDepth+1)
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	next(root, &depth, depth)

	return depth
}

func main() {
	fmt.Println(maxDepth(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
			Right: &TreeNode{
				Val:   7,
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
		}}))
}
