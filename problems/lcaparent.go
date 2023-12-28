package main

func lowestCommonAncestor2(p *TreeNode, q *TreeNode) *TreeNode {

	pPtr, qPtr := p, q

	for pPtr != qPtr {
		if pPtr != nil {
			pPtr = pPtr.Parent
		} else {
			pPtr = q
		}

		if qPtr != nil {
			qPtr = qPtr.Parent
		} else {
			qPtr = p
		}
	}

	return pPtr
}
