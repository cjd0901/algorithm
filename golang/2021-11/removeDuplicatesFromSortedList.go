package _021_11

// Remove Duplicates from Sorted List
// leetcode: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveDuplicatesFromSortedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	virtual := &ListNode{
		Next: head,
	}
	for head.Next != nil {
		ln := head.Next
		if ln.Val == head.Val {
			head.Next = ln.Next
		} else {
			head = ln
		}
	}
	return virtual.Next
}
