package main

import (
	"encoding/json"
	"fmt"
)

// ListNode is the starting data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	switch {
	case l1 == nil && l2 == nil:
		return nil
	case l2 == nil && l1 != nil:
		return l1
	case l1 == nil && l2 != nil:
		return l2
	case (&ListNode{}) == l1 && (&ListNode{}) == l2:
		return &ListNode{}
	case (&ListNode{}) == l1:
		return l2
	case (&ListNode{}) == l2:
		return l1
	}

	fmt.Println(l1.Val, l2.Val)

	switch {
	case l1.Next != nil && l2.Next != nil:
		if l1.Val < l2.Val {
			l3.Val = l1.Val
			l3.Next = mergeTwoLists(l1.Next, l2)
		} else {
			l3.Val = l2.Val
			l3.Next = mergeTwoLists(l2.Next, l1)
		}
	case l1.Next != nil && l2.Next == nil:
		if l1.Val < l2.Val {
			l3.Val = l1.Val
			l3.Next = mergeTwoLists(l1.Next, l2)
		} else {
			l3.Val = l2.Val
			l3.Next = mergeTwoLists(l2.Next, l1)
		}

	case l1.Next == nil && l2.Next != nil:
		if l1.Val < l2.Val {
			l3.Val = l1.Val
			l3.Next = mergeTwoLists(l1.Next, l2)
		} else {
			l3.Val = l2.Val
			l3.Next = mergeTwoLists(l2.Next, l1)
		}

	default:
		if l1.Val < l2.Val {
			l3.Val = l1.Val
			l3.Next = &ListNode{Val: l2.Val}
		} else {
			l3.Val = l2.Val
			l3.Next = &ListNode{Val: l1.Val}
		}
	}
	return l3
}

func (n *ListNode) json() {
	j, err := json.MarshalIndent(n, "", "   ")
	if err != nil {
		fmt.Println("Error printing node json", err)
	}

	fmt.Println(string(j))
}

func main() {
	/*
			node1 := &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
					},
				},
			}

			node2 := &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			}

			finList := mergeTwoLists(node1, node2)
			finList.json()

			node3 := &ListNode{}
			node4 := &ListNode{}
			nList := mergeTwoLists(node3, node4)
			nList.json()

		node5 := &ListNode{
			Val: 2,
		}

		node6 := &ListNode{
			Val: 1,
		}

		xList := mergeTwoLists(node5, node6)
		xList.json()*/

	node1 := &ListNode{
		Val: -9,
		Next: &ListNode{
			Val: 3,
		},
	}

	node2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 7,
		},
	}

	finList := mergeTwoLists(node1, node2)
	finList.json()
}
