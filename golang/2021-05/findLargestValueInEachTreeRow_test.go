package _021_05

import (
	"fmt"
	"testing"
)

func TestFindLargestValueInEachTreeRow(t *testing.T) {
	root := &TreeNode{Val: 1}
	n1 := &TreeNode{Val: 3}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 9}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n4
	n2.Right = n5

	ans := FindLargestValueInEachTreeRow(root)
	fmt.Println(ans)
}