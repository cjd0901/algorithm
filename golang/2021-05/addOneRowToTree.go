package _021_05

import "fmt"

// Add One Row to Tree
// leetcode: https://leetcode-cn.com/problems/add-one-row-to-tree/
// Given the root of a binary tree and two integers val and depth, add a row of nodes with value val at the given depth depth.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func AddOneRowToTree(root *TreeNode, val, depth int) *TreeNode {
	if depth == 1 {
		tn := &TreeNode{Val: val}
		tn.Left = root
		return tn
	}
	queue := []*TreeNode{root}
	d := 1
	for d < depth - 1 {
		temp := []*TreeNode{}
		for len(queue) != 0 {
			tn := queue[0]
			queue = queue[1:]
			if tn.Left != nil {
				temp = append(temp, tn.Left)
			}
			if tn.Right != nil {
				temp = append(temp, tn.Right)
			}
		}
		queue = temp
		d++
	}
	for _, tn := range queue {
		temp := tn.Left
		tn.Left = &TreeNode{Val: val}
		tn.Left.Left = temp
		temp = tn.Right
		tn.Right = &TreeNode{Val: val}
		tn.Right.Right = temp
	}
	return root
}