package _021_07

// Diameter of Binary Tree
// leetcode: https://leetcode-cn.com/problems/diameter-of-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	var dfs func(tn *TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		l := dfs(tn.Left)
		r := dfs(tn.Right)
		diameter = max(diameter, l+r)
		return max(l, r) + 1
	}
	dfs(root)
	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
