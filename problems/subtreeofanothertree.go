package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	res := isSame(root, subRoot)
	if res {
		return true
	}

	leftRes := isSubtree(root.Left, subRoot)
	if leftRes {
		return true
	}
	rightRes := isSubtree(root.Right, subRoot)
	if rightRes {
		return true
	}

	return false
}

func isSame(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	leftRes := isSame(node1.Left, node2.Left)
	rightRes := isSame(node1.Right, node2.Right)

	if leftRes && rightRes {
		return true
	}

	return false
}
