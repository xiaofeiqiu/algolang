package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := make([][]int, 0)

	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]

			level = append(level, curr.Val)

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		res = append(res, level)
	}

	return res
}
