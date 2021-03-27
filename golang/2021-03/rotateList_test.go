package _021_03

import "testing"

func TestRotateList(t *testing.T) {
	head := &ListNode{Val:  1}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 3}
	head.Next = n1
	n1.Next = n2
	RotateList(head, 1)
}
