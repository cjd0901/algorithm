package _021_07

// Balanced Binary Tree
// leetcode: https://leetcode-cn.com/problems/balanced-binary-tree/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func balancedBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left) - height(root.Right)) <= 1 && balancedBinaryTree(root.Left) && balancedBinaryTree(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}