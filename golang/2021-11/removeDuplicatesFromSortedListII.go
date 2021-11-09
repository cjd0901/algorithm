package _021_11

import "fmt"

// Remove Duplicates from Sorted List II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveDuplicatesFromSotredListII(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
