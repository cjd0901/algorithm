package _021_04

// Range Sum of BST
// leetCode: https://leetcode-cn.com/problems/range-sum-of-bst/
// Given the root node of a binary search tree, return the sum of values of all nodes with a value in the range [low, high].

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func RangeSumOfBST(root *TreeNode, low, high int) int {
	sum := 0
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.Val >= low && tn.Val <= high {
			sum += tn.Val
		}
		dfs(tn.Left)
		dfs(tn.Right)
	}
	dfs(root)
	return sum
}

func RangeSumOfBST2(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		RangeSumOfBST2(root.Left, low, high)
	}
	if root.Val < low {
		RangeSumOfBST2(root.Right, low, high)
	}
	return root.Val + RangeSumOfBST2(root.Left, low, high) + RangeSumOfBST2(root.Right, low, high)
}
