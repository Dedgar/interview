/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func preorderTraversal(root *TreeNode) []int {
    var arr []int

    if root == nil {
        return arr
    }

    if (TreeNode{}) == *root {
        return arr
    }


    if root.Left != nil {
        if (&TreeNode{}) != root.Left {
            add := preorderTraversal(root.Left)
            arr = append(arr, add...)
        }
    }

    arr = append(arr, root.Val)

    if root.Right != nil {
        if (&TreeNode{}) != root.Right {
            add := preorderTraversal(root.Right)
            arr = append(arr, add...)
	    }
    }

    return arr
}
