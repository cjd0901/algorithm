package _021_05

import "math"

// Find Largest Value in Each Tree Row
// leetcode: https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/
// Given the root of a binary tree, return an array of the largest value in each row of the tree (0-indexed).

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func FindLargestValueInEachTreeRow(root *TreeNode) []int {
	maxVals := []int{}
	if root == nil {
		return maxVals
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		largest := math.MinInt32
		temp := []*TreeNode{}
		for len(queue) > 0 {
			tn := queue[0]
			queue = queue[1:]
			if tn.Val > largest {
				largest = tn.Val
			}
			if tn.Left != nil {
				temp = append(temp, tn.Left)
			}
			if tn.Right != nil {
				temp = append(temp, tn.Right)
			}
		}
		maxVals = append(maxVals, largest)
		queue = temp
	}
	return maxVals
}