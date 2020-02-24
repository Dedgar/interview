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

func buildTree(inorder []int, postorder []int) *TreeNode {

	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	i := 0

	fmt.Println("rootVal", rootVal)
	for inorder[i] != rootVal {
		fmt.Println("i:", i, "inorder:", inorder[i])
		i++
	}

	// equivalent of a pop operation. not actually removing items in the arrays, just skipping them
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

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
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	fmt.Println("starting with")
	fmt.Println(inorder)
	fmt.Println(postorder)

	aTree := buildTree(inorder, postorder)
	aTree.json()
}
