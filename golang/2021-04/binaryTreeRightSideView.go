package _021_04

import "fmt"

// Binary Tree Right Side View
// leetCode: https://leetcode-cn.com/problems/binary-tree-right-side-view/
// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreeRightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			tn := queue[0]
			if i == l-1 {
				res = append(res, tn.Val)
			}
			if tn.Left != nil {
				queue = append(queue, tn.Left)
			}
			if tn.Right != nil {
				queue = append(queue, tn.Right)
			}
			queue = queue[1:]
		}
	}
	return res
}
