package _021_07

// Second Minimum Node In a Binary Tree
// leetcode:https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func secondMinimumNodeInABinaryTree(root *TreeNode) int {
	min, ans := root.Val, -1
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.Val > min && (tn.Val < ans || ans == -1) {
			ans = tn.Val
		}
		if tn.Val < min {
			ans = min
			min = tn.Val
		}
		dfs(tn.Left)
		dfs(tn.Right)
	}
	dfs(root)
	return ans
}