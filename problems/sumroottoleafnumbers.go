package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	res := 0
	dfs13(root, 0, &res)
	return res
}

func dfs13(root *TreeNode, num int, res *int) {
	if root == nil {
		return
	}

	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += num
		return
	}
	dfs13(root.Left, num, res)
	dfs13(root.Right, num, res)
}
