package _021_04

import (
	"fmt"
	"testing"
)

func TestRangeSumOfBST(t *testing.T) {
	n1 := &TreeNode{Val: 10}
	n2 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 15}
	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 7}
	n6 := &TreeNode{Val: 18}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Right = n6
	sum := RangeSumOfBST(n1, 7, 15)
	fmt.Println(sum)
}