package _021_07

import (
	"fmt"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 5}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	count := DiameterOfBinaryTree(root)
	fmt.Println(count)
}