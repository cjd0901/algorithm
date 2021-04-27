package _021_04

import (
	"fmt"
	"testing"
)

func TestNaryTreeLevelOrderTraversal(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 3}
	n3 := &Node{Val: 2}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n1.Children = []*Node{n2, n3, n4}
	n2.Children = []*Node{n5, n6}
	ans := NaryTreeLevelOrderTraversal(n1)
	fmt.Println(ans)
}