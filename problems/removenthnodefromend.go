package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := dummy, dummy

	for i := 0; i < n+1; i++ {
		second = second.Next
	}

	for second != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = first.Next.Next
	return dummy.Next
}
