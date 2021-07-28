package _021_07

// All Nodes Distance K in Binary Tree
// leetcode: https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func AllNodesDistanceKInBinaryTree(root *TreeNode, target *TreeNode, k int) []int {
	parents := make(map[int]*TreeNode)
	var addParents func(tn *TreeNode)
	addParents = func(tn *TreeNode) {
		if tn.Left != nil {
			parents[tn.Left.Val] = tn
			addParents(tn.Left)
		}
		if tn.Right != nil {
			parents[tn.Right.Val] = tn
			addParents(tn.Right)
		}
	}
	addParents(root)
	var ans []int
	var dfs func(tn, from *TreeNode, depth int)
	dfs = func(tn, from *TreeNode, depth int) {
		if tn == nil {
			return
		}
		if depth == k {
			ans = append(ans, tn.Val)
			return
		}
		if tn.Left != from {
			dfs(tn.Left, tn, depth+1)
		}
		if tn.Right != from {
			dfs(tn.Right, tn, depth+1)
		}
		if parents[tn.Val] != from {
			dfs(parents[tn.Val], tn, depth+1)
		}
	}
	dfs(target, nil, 0)
	return ans
}