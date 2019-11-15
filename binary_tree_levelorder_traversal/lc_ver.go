package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// we're appending to the array inside the main array by index.
// lvl 3 of tree = append(arr[3], somenum), where lvl 0 is the root node.
func next(root *TreeNode, depth int, arr *[][]int) {
	if root.Left != nil {
		if (&TreeNode{}) != root.Left {
			lVal := root.Left.Val
			//fmt.Println("left val is ", lVal)
			if len(*arr) > depth {
				(*arr)[depth] = append((*arr)[depth], lVal)
			} else {
				(*arr) = append(*arr, []int{lVal})
			}
			next(root.Left, depth+1, arr)
		}
	}

	if root.Right != nil {
		if (&TreeNode{}) != root.Right {
			rVal := root.Right.Val
			//fmt.Println("right val is ", rVal)
			if len(*arr) > depth {
				(*arr)[depth] = append((*arr)[depth], rVal)
			} else {
				(*arr) = append(*arr, []int{rVal})
			}
			next(root.Right, depth+1, arr)
		}
	}
}

func levelOrder(root *TreeNode) [][]int {
	var arr [][]int

	if root == nil {
		return arr
	}

	if (TreeNode{}) == *root {
		return arr
	}

	arr = append(arr, []int{root.Val})

	next(root, 1, &arr)

	return arr
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
	fmt.Println(levelOrder(&TreeNode{
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
