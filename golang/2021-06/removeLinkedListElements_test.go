package _021_06

import (
	"fmt"
	"testing"
)

func TestRemoveLinkedListElements(t *testing.T) {
	head := &ListNode{Val: 7}
	ln1 := &ListNode{Val: 7}
	ln2 := &ListNode{Val: 7}
	ln3 := &ListNode{Val: 7}
	head.Next = ln1
	ln1.Next = ln2
	ln2.Next = ln3
	ln := RemoveLinkedListElements2(head, 7)
	for ln != nil {
		fmt.Println(ln.Val)
		ln = ln.Next
	}
}