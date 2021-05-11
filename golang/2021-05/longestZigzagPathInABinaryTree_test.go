package _021_05

import (
	"fmt"
	"testing"
)

func TestLongestZigzagPathInABinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 1}
	root.Left = n1
	n1.Right = n2
	path := LongestZigzagPathInABinaryTree(root)
	fmt.Println(path)
}