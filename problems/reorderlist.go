package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first, second := head, prev
	// why check second.Next instead of second != nil ?
	// this is because the first half might be longer than the second half
	// when second half reaches to the end, first half still have 1 node left
	// and no need to move the last node
	// so we can exit the loop ealier to prevent accseesing nil pointer in the second list
	for second.Next != nil {
		firstNext, secondNext := first.Next, second.Next
		first.Next = second
		second.Next = firstNext
		first, second = firstNext, secondNext
	}
}
