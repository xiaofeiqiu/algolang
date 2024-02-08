package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}

	beforeStart := dummy
	for i := 0; i < left-1; i++ {
		beforeStart = beforeStart.Next
	}

	// the start node will be the end node after reverse
	// so store the start node
	start := beforeStart.Next

	curr := start
	var prev *ListNode

	for i := 0; i < right-left+1; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// connect reversed tailed to the original list
	start.Next = curr

	// connect original list to the head of reversed node
	beforeStart.Next = prev

	return dummy.Next
}
