package _021_05

// Leaf-similar Trees
// leetcode: https://leetcode-cn.com/problems/leaf-similar-trees/
// Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func LeafSimilarTrees(root1, root2 *TreeNode) bool {
	arr := []int{}
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.Left == nil && tn.Right == nil {
			arr = append(arr, tn.Val)
			return
		}
		dfs(tn.Left)
		dfs(tn.Right)
	}
	dfs(root1)
	arr1 := append([]int(nil), arr...)
	arr = []int{}
	dfs(root2)
	if len(arr) != len(arr1) {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr1[i] {
			return false
		}
	}
	return true
}