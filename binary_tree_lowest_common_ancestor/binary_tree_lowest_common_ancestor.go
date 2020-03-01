package main

import (
	"encoding/json"
	"fmt"
)

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root == p || root == q:
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	switch {
	case left != nil && right != nil:
		return root
	case left == nil:
		return right
	case right == nil:
		return left
	}

	return nil
}

func (n *TreeNode) json() {
	j, err := json.MarshalIndent(n, "", "   ")
	if err != nil {
		fmt.Println("Error printing node json", err)
	}

	fmt.Println(string(j))
}

func main() {
	root := &TreeNode{
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
		}}

	p := &TreeNode{
		Val:   15,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	q := &TreeNode{
		Val:   7,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}

	result := lowestCommonAncestor(root, p, q)

	result.json()
}
