package main

func kthSmallest(root *TreeNode, k int) int {
	res := 0
	dfs1(root, &res, &k)

	return res
}

func dfs1(curr *TreeNode, res *int, k *int) {
	if curr == nil {
		return
	}

	dfs1(curr.Left, res, k)
	*k--
	if *k == 0 {
		*res = curr.Val
		return
	}
	dfs1(curr.Right, res, k)
}
