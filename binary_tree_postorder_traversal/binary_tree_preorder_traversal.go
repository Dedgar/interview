package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) checkTree(arr *[]int) {
	fmt.Println("new instance")
	*arr = append(*arr, root.Val)

	fmt.Println("check left")
	if *root.Left != (TreeNode{}) {
		root.Left.checkTree(arr)
	}

	fmt.Println("check right")
	if *root.Right != (TreeNode{}) {
		root.Right.checkTree(arr)
	}
}

func preorderTraversal(root *TreeNode) []int {
	var arr []int

	root.checkTree(&arr)

	return arr
}

func main() {
	fmt.Println(preorderTraversal(&TreeNode{Val: 1, Left: &TreeNode{}, Right: &TreeNode{Val: 2, Left: &TreeNode{}, Right: &TreeNode{}}}))
}
