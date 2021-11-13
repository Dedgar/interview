package main

import "fmt"

// ListNode defines a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode finds either the middle node of a singly-linked list, or the second middle node for even-numbered lists
// a straightforward way to find the middle of singly linked list is to skip steps on the way, while still keeping track of the overall steps.
// half of whatever our skip-to-the-end pointer is will to equal our current iteration step.
//
// 1   3   5   7   9    11    13    15
// t   t   t   t   t    t     t     t
// b b b b b b b b b
// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
func middleNode(head *ListNode) *ListNode {
	bot := head
	top := head

	for top != nil && top.Next != nil {
		bot = bot.Next
		top = top.Next.Next
	}
	return bot
}

func main() {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	node2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	fmt.Println(*middleNode(node1))
	fmt.Println(*middleNode(node2))
}
