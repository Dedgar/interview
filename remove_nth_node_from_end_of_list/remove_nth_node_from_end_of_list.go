package main

// ListNode defines a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// new definition of head to offset the data structure length.
	// allows us to simplify our algorithm and use a single for loop referring to .Next items.
	head = &ListNode{Next: head}
	bot := head
	top := head

	// iterate until there are no more non-nil .Next values.
	// when we reach nil, we're at the end of the singly-linked list.
	for pos := 0; bot != nil; pos++ {
		top = top.Next
		bot = bot.Next

		if pos == n {
			top = head
		}
	}

	// finally, we now know we're at the point in the list matching n.
	// replace n.Next with n.Next.Next. works even if its value is nil and it's already at the end of the list.
	top.Next = top.Next.Next

	return head.Next
}

/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fmt.Println("head is", head)
	bot := head
	top := head

	for pos := 0; top != nil; pos++ {

		bot = bot.Next
		top = top.Next
		if pos == n {
			bot = head
		}
	}
	return head.Next
}
*/
/*
		//		for top != nil && top.Next != nil {
		//			top = top.Next.Next
		//		}
func printHead(ln ListNode) {
	top := ln

	for top.Next != nil {
		fmt.Println(ln.Val)
		top = top.Next.Next
	}
}
*/

func main() {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
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

	removeNthFromEnd(node1, 1)
	removeNthFromEnd(node2, 3)

	//	printHead(*node1)
	//	printHead(*node2)
}
