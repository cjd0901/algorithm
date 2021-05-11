package _021_05

// Print Binary Tree From Top to Bottom
// leetcode: https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func PrintBinaryTreeFromTopToBottom(root *TreeNode) []int {
	prints, queue := []int{}, []*TreeNode{root}
	if root == nil {
		return prints
	}
	for len(queue) > 0 {
		temp := []*TreeNode{}
		for len(queue) > 0 {
			tn := queue[0]
			queue = queue[1:]
			prints = append(prints, tn.Val)
			if tn.Left != nil {
				temp = append(temp, tn.Left)
			}
			if tn.Right != nil {
				temp = append(temp, tn.Right)
			}
		}
		queue = temp
	}
	return prints
}