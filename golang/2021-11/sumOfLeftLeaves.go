package _021_11

// Sum of Left Leaves
// leetcode: https://leetcode-cn.com/problems/sum-of-left-leaves/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func SumOfLeftLeaves(root *TreeNode) int {
	ans := 0
	var dfs func(tn *TreeNode, left bool)
	dfs = func(tn *TreeNode, left bool) {
		if tn == nil {
			return
		}
		if left && tn.Left == nil && tn.Right == nil {
			ans += tn.Val
			return
		}
		dfs(tn.Left, true)
		dfs(tn.Right, false)
	}
	dfs(root.Left, true)
	dfs(root.Right, false)
	return ans
}

func SumOfLeftLeaves2(root *TreeNode) int {
	var dfs func(tn *TreeNode) int
	dfs = func(tn *TreeNode) int {
		ans := 0
		if tn.Left != nil {
			if IsLeave(tn.Left) {
				ans += tn.Left.Val
			} else {
				ans += dfs(tn.Left)
			}
		}
		if tn.Right != nil && !IsLeave(tn.Right) {
			ans += dfs(tn.Right)
		}
		return ans
	}
	return dfs(root)
}

func IsLeave(tn *TreeNode) bool {
	return tn.Left == nil && tn.Right == nil
}