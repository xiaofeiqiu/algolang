package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return isValid1(root, math.MinInt, math.MaxInt)
}

func isValid1(node *TreeNode, low int, high int) bool {
	if node == nil {
		return true
	}

	// check left subtree, update high to current node's val
	left := isValid1(node.Left, low, node.Val)

	// check right subtree, update low to current list' val
	right := isValid1(node.Right, node.Val, high)

	// check if current root node is in the valid range
	if node.Val <= low || node.Val >= high {
		return false
	}

	return left && right
}
