package _021_04

import (
	"fmt"
	"testing"
)

func TestBinaryTreeRightSideView(t *testing.T) {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 2, Left: nil}
	n2 := &TreeNode{Val: 3, Left: nil}
	n3 := &TreeNode{Val: 4, Left: nil, Right: nil}
	n4 := &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left = n1
	root.Right = n2
	n1.Right = n4
	n2.Right = n3
	view := BinaryTreeRightSideView(root)
	fmt.Println(view)
}