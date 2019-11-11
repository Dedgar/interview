package main

import (
	"encoding/json"
	"fmt"
)

// ListNode is the definition for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head {
		fmt.Println("head is nil")
		return head
	}

	if (ListNode{}) == *head {
		fmt.Println("returning early")
		return head
	}

	fmt.Println("current", head.Val)

	if nil != head.Next {
		if head.Val == head.Next.Val {
			fmt.Printf("%v does equal %v, replacing it with next %v\n", head.Val, head.Next.Val, head.Next.Val)
			*head = *head.Next
			// now that we've skipped ahead a node, calling deleteDuplicates(head) in this manner is
			// the equivalent of deleting a node and skipping ahead with deleteDuplicates(head.Next)
			deleteDuplicates(head)
		} else {
			// skip this node
			deleteDuplicates(head.Next)
		}
	}

	return head
}

func (n *ListNode) json() {
	j, err := json.MarshalIndent(n, "", "   ")
	if err != nil {
		fmt.Println("Error printing node json", err)
	}
	fmt.Println(string(j))
}

func main() {
	newList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 2,
							},
						},
					},
				},
			},
		},
	}
	fmt.Println("starting with")
	newList.json()

	deleteDuplicates(newList)

	fmt.Println("ending with")
	newList.json()
}
