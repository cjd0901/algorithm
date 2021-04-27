package _021_04

// N-ary Tree Level Order Traversal
// leetCode: https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
// Given an n-ary tree, return the level order traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See examples).

type Node struct {
	Val int
	Children []*Node
}

func NaryTreeLevelOrderTraversal(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		temp := []int{}
		tempQ := queue
		queue = queue[len(queue):]
		for _, n := range tempQ {
			temp = append(temp, n.Val)
			if len(n.Children) > 0 {
				queue = append(queue, n.Children...)
			}
		}
		ans = append(ans, temp)
	}
	return ans
}