package _021_04

// Increasing Order Search Tree
// leetCode: https://leetcode-cn.com/problems/increasing-order-search-tree/
// Given the root of a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only one right child.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IncreasingOrderSearchTree(root *TreeNode) *TreeNode {
	var n, ans *TreeNode
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		if n == nil {
			n = &TreeNode{
				Val: tn.Val,
			}
			ans = n
		}else {
			n.Right = &TreeNode{
				Val: tn.Val,
			}
			n = n.Right
		}
		dfs(tn.Right)
	}
	dfs(root)
	return ans
}
