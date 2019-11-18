package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func next(root1, root2 *TreeNode) bool {
	switch {
	case root1 == nil && root2 == nil:
		return true
	case root1 == nil || root2 == nil:
		return false
	default:
		return (root1.Val == root2.Val) &&
			next(root1.Right, root2.Left) &&
			next(root1.Left, root2.Right)
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if (TreeNode{}) == *root {
		return true
	}

	return next(root, root)
}

func main() {
	/*
		fmt.Println(levelOrder(&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
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
	*/
	fmt.Println(isSymmetric(&TreeNode{
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
