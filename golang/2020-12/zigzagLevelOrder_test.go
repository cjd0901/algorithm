package golang

import (
	"fmt"
	"testing"
)

func TestZigZagLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val:   3,
	}
	n1 := &TreeNode{
		Val: 9,
	}
	n2 := &TreeNode{
		Val: 20,
	}
	n3 := &TreeNode{
		Val: 15,
	}
	n4 := &TreeNode{
		Val: 7,
	}
	root.Left = n1
	root.Right = n2
	n2.Left = n3
	n2.Right = n4
	order := ZigzagLevelOrder2(root)
	fmt.Println(order)
}