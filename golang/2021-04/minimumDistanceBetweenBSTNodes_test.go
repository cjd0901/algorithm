package _021_04

import (
	"fmt"
	"testing"
)

func TestMinimumDistanceBetweenBSTNodes(t *testing.T) {
	root := TreeNode{Val: 1}
	n1 := TreeNode{Val: 0, Left: nil, Right: nil}
	n2 := TreeNode{Val: 48}
	n3 := TreeNode{Val: 12, Left: nil, Right: nil}
	n4 := TreeNode{Val: 49, Left: nil, Right: nil}
	root.Left = &n1
	root.Right = &n2
	n1.Left = &n3
	n1.Right = &n4
	ans := MinimumDistanceBetweenBSTNodes(&root)
	fmt.Println(ans)
}
