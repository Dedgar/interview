package main

import (
	"fmt"
	"strconv"
)

// TreeNode is tghe definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func search(root *TreeNode, res *[]string, path string) {
	switch {
	case root.Left == nil && root.Right == nil:
		*res = append(*res, path)
		return

	case root.Left != nil:
		search(root.Left, res, path+"->"+strconv.Itoa(root.Left.Val))

	case root.Right != nil:
		search(root.Right, res, path+"->"+strconv.Itoa(root.Right.Val))
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string

	if root == nil {
		return res
	}
	search(root, &res, strconv.Itoa(root.Val))
	return res
}

func main() {
	fmt.Println(binaryTreePaths(&TreeNode{
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
