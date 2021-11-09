package _021_11

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesFromSotredListII(t *testing.T) {
	head := &ListNode{Val: 1}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 4}
	n6 := &ListNode{Val: 5}
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	ans := RemoveDuplicatesFromSotredListII(head)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}