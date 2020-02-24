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

func checkNodes(found, root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}

	/*
		if checkNodes(found, root.Left, p, q) || checkNodes(found, root.Right, p, q) {
			*found = *root
			fmt.Println("child val matches", root.Val, p.Val, q.Val)
			return true
		}
	*/

	if root.Val == p.Val || root.Val == q.Val {
		*found = *root
		fmt.Println("root val matches", root.Val, p.Val, q.Val)
		return true
	}

	return false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	found := &TreeNode{}

	checkNodes(found, root, p, q)

	return found
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
