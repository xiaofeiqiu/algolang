package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}

	// if node > key, delete key in left subtree and update root.left
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	// if node < key, delete key in right subtree and update root.right
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	// if root has no child, remove node (return nil)
	if root.Left == nil && root.Right == nil {
		return nil
	}

	// if root has right child, return right subtree
	if root.Left == nil {
		return root.Right
	}

	// if root has left child, return left subtree
	if root.Right == nil {
		return root.Left
	}

	// if root has 2 children, fin predecessor
	pre := findPredecessor(root)

	// replce root with the predecessor
	root.Val = pre.Val

	// remove predecessor in left subtree
	root.Left = deleteNode(root.Left, pre.Val)
	return root
}

// findPredecessor, predecessor is in the most right of the left subtree
func findPredecessor(root *TreeNode) *TreeNode {
	node := root.Left
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// successor is in the most left of the right subtree
func findSuccessor(root *TreeNode) *TreeNode {
	node := root.Right
	for node.Left != nil {
		node = node.Left
	}
	return node
}
