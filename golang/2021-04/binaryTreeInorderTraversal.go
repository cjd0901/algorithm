package _021_04

// Binary Tree Inorder Traversal
// leetcode: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// Given the root of a binary tree, return the inorder traversal of its nodes' values.

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func BinaryTreeInorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var inorder func(tn *TreeNode)
	inorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inorder(tn.Left)
		ans = append(ans, tn.Val)
		inorder(tn.Right)
	}
	inorder(root)
	return ans
}