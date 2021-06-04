package _021_06

import (
	"fmt"
	"testing"
)

func TestIntersectionOfTwoLinkedLists(t *testing.T) {
	headA := &ListNode{Val: 4}
	headA2 := &ListNode{Val: 1}
	headB := &ListNode{Val: 5}
	headB2 := &ListNode{Val: 6}
	headB3 := &ListNode{Val: 1}
	c1 := &ListNode{Val: 8}
	c2 := &ListNode{Val: 4}
	c3 := &ListNode{Val: 5}
	c1.Next = c2
	c2.Next = c3
	headA.Next = headA2
	headA2.Next = c1
	headB.Next = headB2
	headB2.Next = headB3
	headB3.Next = c1
	tn := IntersectionOfTwoLinkedLists(headA, headB)
	fmt.Println(tn.Val)
}