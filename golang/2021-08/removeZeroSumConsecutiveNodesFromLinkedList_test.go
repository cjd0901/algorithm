package _021_08

import (
	"fmt"
	"testing"
)

func TestRemoveZeroSumSublists(t *testing.T) {
	head := &ListNode{Val: 1}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: -3}
	n4 := &ListNode{Val: 4}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	ans := RemoveZeroSumSublists(head)
	fmt.Println(ans)
}