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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	i := 0

	fmt.Println("rootVal", rootVal)
	for i < len(inorder) {
		if inorder[i] == rootVal {
			fmt.Println("i:", i, "inorder:", inorder[i])
			break
		}
		i++
	}

	// equivalent of a pop operation. not actually removing items in the arrays, just skipping them
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}

func (n *TreeNode) json() {
	j, err := json.MarshalIndent(n, "", "   ")
	if err != nil {
		fmt.Println("Error printing node json", err)
	}

	fmt.Println(string(j))
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	fmt.Println("starting with")
	fmt.Println(preorder)
	fmt.Println(inorder)

	aTree := buildTree(preorder, inorder)
	aTree.json()
}
