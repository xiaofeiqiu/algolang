package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestorbst(root, p, q *TreeNode) *TreeNode {

	for root != nil {
		if root.Val < q.Val && root.Val < p.Val {
			root = root.Right
			continue
		}

		if root.Val > q.Val && root.Val > p.Val {
			root = root.Left
			continue
		}

		return root
	}

	return nil
}
