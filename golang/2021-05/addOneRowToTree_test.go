package _021_05

import "testing"

func TestAddOneRowToTree(t *testing.T) {
	root := &TreeNode{Val: 4}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 6}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 1}
	n5 := &TreeNode{Val: 5}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Right = n5
	AddOneRowToTree(root, 1, 2)
}