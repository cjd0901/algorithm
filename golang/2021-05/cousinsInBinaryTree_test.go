package _021_05

import (
	"fmt"
	"testing"
)

func TestCousinsInBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	b := CousinsInBinaryTree(root, 2, 3)
	fmt.Println(b)
}