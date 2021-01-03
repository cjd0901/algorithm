package _021_01

import (
	"fmt"
	"testing"
)

func TestPartitionList(t *testing.T) {
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 5}
	node5 := &ListNode{Val: 2}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	list := PartitionList(head, 3)
	fmt.Println(list.Val)
}