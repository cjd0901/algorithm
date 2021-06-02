package _021_06

import (
	"fmt"
	"testing"
)

func TestReverseLinkedListII(t *testing.T) {
	root := &ListNode{Val: 1}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 4}
	n4 := &ListNode{Val: 5}
	root.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	head := ReverseLinkedListII(root, 2, 4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}