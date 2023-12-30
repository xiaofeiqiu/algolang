package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return helper(lists)
}

func helper(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	left := helper(lists[:mid])
	right := helper(lists[mid:])

	return merge1(left, right)
}

func merge1(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
