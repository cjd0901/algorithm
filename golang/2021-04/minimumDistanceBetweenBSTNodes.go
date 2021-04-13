package _021_04

import (
	"math"
)

// Minimum Distance Between BST Nodes
// leetCode: https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/
// Given the root of a Binary Search Tree (BST), return the minimum difference between the values of any two different nodes in the tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 父子节点直接的最小值
func MinimumDistanceBetweenBSTNodes(root *TreeNode) int {
	distance := math.MaxInt32
	pre := -1
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		if pre != -1 && minus(tn.Val, pre) < distance {
			distance = minus(tn.Val, pre)
		}
		pre = tn.Val
		dfs(tn.Right)
	}
	dfs(root)
	return distance
}

func minus(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
