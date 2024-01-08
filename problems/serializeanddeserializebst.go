/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import (
	"math"
	"strconv"
	"strings"
)

type Codec3 struct {
}

func Constructor4() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize1(root *TreeNode) string {
	res := make([]string, 0)

	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		valStr := strconv.Itoa(node.Val)
		res = append(res, valStr)
		helper(node.Left)
		helper(node.Right)
	}

	helper(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize1(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}

	strs := strings.Split(data, ",")

	var buildTree func(*[]string, int, int) *TreeNode
	buildTree = func(strs *[]string, minBound, maxBound int) *TreeNode {

		if len(*strs) == 0 {
			return nil
		}
		val, _ := strconv.Atoi((*strs)[0])
		if val < minBound || val > maxBound {
			return nil
		}

		*strs = (*strs)[1:]
		node := &TreeNode{Val: val}
		node.Left = buildTree(strs, minBound, val)
		node.Right = buildTree(strs, val, maxBound)
		return node
	}

	return buildTree(&strs, math.MinInt, math.MaxInt)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
