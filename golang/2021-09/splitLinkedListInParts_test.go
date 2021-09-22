package _021_09

import (
	"testing"
)

func TestSplitLinkedListInParts(t *testing.T) {
	root := &ListNode{Val: 0}
	n1 := &ListNode{Val: 2}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 6}
	n4 := &ListNode{Val: 6}
	root.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	ans := SplitLinkedListInParts(root, 3)
	t.Log(ans)
}