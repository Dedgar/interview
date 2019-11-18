package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func next(root *TreeNode, curNum, total int) bool {
	if root == nil {
		return false
	}

	newNum := curNum + root.Val

	if newNum == total {
		if root.Left == nil && root.Right == nil {
			return true
		}
	}

	return next(root.Left, newNum, total) || next(root.Right, newNum, total)
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if (TreeNode{}) == *root {
		return false
	}

	return next(root, 0, sum)
}

func main() {
	fmt.Println(hasPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{},
			Right: &TreeNode{},
		}}, 1))

	/*
		fmt.Println(hasPathSum(&TreeNode{
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
			}}, 36))*/
}
