package main

func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	dummy := &TreeNode{}
	listCurr := dummy
	dfs12(root, &listCurr)

	listCurr.Right = dummy.Right
	dummy.Right.Left = listCurr
	return dummy.Right
}

func dfs12(treeCurr *TreeNode, listCurr **TreeNode) {
	if treeCurr == nil {
		return
	}

	dfs12(treeCurr.Left, listCurr)
	(*listCurr).Right = treeCurr
	treeCurr.Left = *listCurr
	*listCurr = (*listCurr).Right

	dfs12(treeCurr.Right, listCurr)
}
