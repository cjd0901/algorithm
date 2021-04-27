package _021_04

import (
	"fmt"
	"testing"
)

func TestBinaryTreeInorderTraversal(t *testing.T) {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Right = n2
	n2.Left = n3
	ans := BinaryTreeInorderTraversal(n1)
	fmt.Println(ans)
}