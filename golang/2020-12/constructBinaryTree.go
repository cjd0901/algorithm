package golang

import "fmt"

// Construct Binary Tree from Preorder and Inorder Traversal
// leetCode: https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// Given preorder and inorder traversal of a tree, construct the binary tree.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = BuildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = BuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}