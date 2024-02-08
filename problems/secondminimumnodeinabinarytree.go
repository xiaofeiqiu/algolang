package main

import "math"

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	secondMin := findSecondMin(root, root.Val)
	if secondMin == math.MaxInt32 {
		return -1
	}
	return secondMin
}

func findSecondMin(node *TreeNode, firstMin int) int {
	if node == nil {
		return math.MaxInt32
	}
	if node.Val > firstMin {
		return node.Val
	}
	leftSecondMin := findSecondMin(node.Left, firstMin)
	rightSecondMin := findSecondMin(node.Right, firstMin)
	return min(leftSecondMin, rightSecondMin)
}
