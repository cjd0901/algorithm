package _021_05

import "math"

// Maximum Level Sum of a Binary Tree
// leetcode: https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree/
// Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
// Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func MaximumLevelSumOfABinaryTree(root *TreeNode) int {
	maxSum, maxIndex := math.MinInt32, 0
	d := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := []*TreeNode{}
		tempSum := 0
		for len(queue) > 0 {
			tn := queue[0]
			queue = queue[1:]
			tempSum += tn.Val
			if tn.Left != nil {
				temp = append(temp, tn.Left)
			}
			if tn.Right != nil {
				temp = append(temp, tn.Right)
			}
		}
		if tempSum > maxSum {
			maxSum = tempSum
			maxIndex = d
		}
		queue = temp
		d++
	}
	return maxIndex
}