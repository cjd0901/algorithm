package _021_05

import (
	"fmt"
	"testing"
)

func TestLeafSimilarTrees(t *testing.T) {
	root1 := &TreeNode{val: 3}
	tn1 := &TreeNode{val: 5}
	tn2 := &TreeNode{val: 1}
	tn3 := &TreeNode{val: 6}
	tn4 := &TreeNode{val: 2}
	tn5 := &TreeNode{val: 9}
	tn6 := &TreeNode{val: 8}
	tn7 := &TreeNode{val: 7}
	tn8 := &TreeNode{val: 4}
	root1.Left = tn1
	root1.Right = tn2
	tn1.Left = tn3
	tn1.Right = tn4
	tn2.Left = tn5
	tn2.Right = tn6
	tn4.Left = tn7
	tn4.Right = tn8

	similar := LeafSimilarTrees(root1, nil)
	fmt.Println(similar)
}