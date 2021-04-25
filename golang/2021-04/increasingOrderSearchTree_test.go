package _021_04

import (
	"fmt"
	"testing"
)

func TestIncreasingOrderSearchTree(t *testing.T) {
	n1 := &TreeNode{Val: 5}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 6}
	n4 := &TreeNode{Val: 2}
	n5 := &TreeNode{Val: 4}
	n6 := &TreeNode{Val: 8}
	n7 := &TreeNode{Val: 1}
	n8 := &TreeNode{Val: 7}
	n9 := &TreeNode{Val: 9}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Right = n6
	n4.Left = n7
	n6.Left = n8
	n6.Right = n9
	node := IncreasingOrderSearchTree(n1)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Right
	}
}
