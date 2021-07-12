package _021_07

// Paths with Sum
// leetcode: https://leetcode-cn.com/problems/paths-with-sum-lcci/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PathsWithSum(root *TreeNode, sum int) int {
	count := 0
	var dfs func(tn *TreeNode, s int); dfs = func(tn *TreeNode, s int) {
		if tn == nil {
			return
		}
		s -= tn.Val
		if s == 0 {
			count ++
		}
		dfs(tn.Left, s)
		dfs(tn.Right, s)
	}
	var dfs2 func(tn *TreeNode, s int); dfs2 = func(tn *TreeNode, s int) {
		if tn == nil {
			return
		}
		dfs(tn, s)
		dfs2(tn.Left, s)
		dfs2(tn.Right, s)
	}
	dfs2(root, sum)
	return count
}