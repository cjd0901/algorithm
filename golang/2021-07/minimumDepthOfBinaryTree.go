package _021_07

import "math"

// Minimum Depth of Binary Tree
// leetcode: https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func minimumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth := 100000
	var dfs func(tn *TreeNode, depth int)
	dfs = func(tn *TreeNode, depth int) {
		if tn == nil {
			return
		}
		if tn.Right == nil && tn.Left == nil {
			minDepth = min(minDepth, depth)
		}
		dfs(tn.Right, depth + 1)
		dfs(tn.Left, depth + 1)
	}
	dfs(root, 1)
	return minDepth
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}