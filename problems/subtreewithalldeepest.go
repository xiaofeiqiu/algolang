package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, res := dfs4(root)
	return res
}

func dfs4(node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}

	leftDepth, leftResult := dfs4(node.Left)
	rightDepth, rightResult := dfs4(node.Right)

	if leftDepth == rightDepth {
		return leftDepth + 1, node
	}

	if leftDepth > rightDepth {
		return leftDepth + 1, leftResult
	}

	return rightDepth + 1, rightResult
}
