package _021_05

// Longest Zigzag Path in a Binary Tree
// leetcode: https://leetcode-cn.com/problems/longest-zigzag-path-in-a-binary-tree/
// You are given the root of a binary tree.
// A ZigZag path for a binary tree is defined as follow:
// Choose any node in the binary tree and a direction (right or left).
// If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
// Change the direction from right to left or from left to right.
// Repeat the second and third steps until you can't move in the tree.
// Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
// Return the longest ZigZag path contained in that tree.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LongestZigzagPathInABinaryTree(root *TreeNode) int {
	longestPath := 0
	if root == nil {
		return longestPath
	}
	// fDir true 为父节点是左节点
	var dfs func(tn *TreeNode, fDir bool, l int)
	dfs = func(tn *TreeNode, fDir bool, l int) {
		longestPath = max(longestPath, l)
		if fDir {
			if tn.Left != nil {
				dfs(tn.Left, true, 1)
			}
			if tn.Right != nil {
				dfs(tn.Right, false, l+1)
			}
		}else {
			if tn.Left != nil {
				dfs(tn.Left, true, l+1)
			}
			if tn.Right != nil {
				dfs(tn.Right, false, 1)
			}
		}
	}
	dfs(root, true, 0)
	dfs(root, false, 0)
	return longestPath
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}