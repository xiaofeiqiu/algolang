package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

type Node struct {
	Val       int
	Neighbors []*Node
}
