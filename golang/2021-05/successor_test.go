package _021_05

import (
	"fmt"
	"testing"
)

func TestSuccessor(t *testing.T) {
	root := &TreeNode{Val: 2}
	tr1 := &TreeNode{Val: 1}
	tr2 := &TreeNode{Val: 3}
	root.Left = tr1
	root.Right = tr2
	successor := Successor(root, tr1)
	fmt.Println(successor)
}