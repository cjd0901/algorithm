package golang

// Binary Tree Zigzag Level Order Traversal
// leetCode: https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
// Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 数组
func ZigzagLevelOrder(root *TreeNode) (L [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		q := queue
		queue = nil
		values := []int{}
		for _, node := range q {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			values = append(values, node.Val)
		}
		if level%2 == 1 {
			for i := 0; i < len(values) / 2; i++ {
				values[i], values[len(values) - i - 1] = values[len(values) - i -1], values[i]
			}
		}
		L = append(L, values)
	}
	return
}

// 队列
func ZigzagLevelOrder2(root *TreeNode) (L [][]int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	isLeftStart := true
	for len(queue) > 0 {
		l := len(queue)
		nodes := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeftStart {
				nodes[i] = node.Val
			}else {
				nodes[l-1-i] = node.Val
			}
		}
		L = append(L, nodes)
		isLeftStart = !isLeftStart
		queue = queue[l:]
	}
	return
}