package _021_07

import (
	"fmt"
	"testing"
)

func TestPathsWithSum(t *testing.T) {
	root := &TreeNode{Val: 1}
	tn1 := &TreeNode{Val: 2}
	tn2 := &TreeNode{Val: 3}
	tn3 := &TreeNode{Val: 4}
	tn4 := &TreeNode{Val: 5}
	root.Right = tn1
	tn1.Right = tn2
	tn2.Right = tn3
	tn3.Right = tn4
	count := PathsWithSum(root, 3)
	fmt.Println(count)
}