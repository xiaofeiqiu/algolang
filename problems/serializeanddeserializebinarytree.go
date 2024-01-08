package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec1 struct {
}

func Constructor3() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	res := make([]string, 0)
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			res = append(res, "null")
			return
		}

		str := strconv.Itoa(node.Val)
		res = append(res, str)

		helper(node.Left)
		helper(node.Right)
	}

	helper(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
	index := 0

	var helper func([]string) *TreeNode
	helper = func(strs []string) *TreeNode {
		if strs[index] == "null" {
			index++
			return nil
		}

		val, err := strconv.Atoi(strs[index])
		if err != nil {
			return nil
		}
		index++

		return &TreeNode{
			Val:   val,
			Left:  helper(strs),
			Right: helper(strs),
		}
	}

	return helper(strs)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
