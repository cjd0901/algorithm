package _021_05

import (
	"fmt"
	"testing"
)

func TestLinkedListComponents(t *testing.T) {
	head := &ListNode{Val: 0}
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	head.Next = l1
	l1.Next = l2
	l2.Next = l3
	ans := LinkedListComponents(head, []int{0, 1, 3})
	fmt.Println(ans)
}