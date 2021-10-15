package _021_10

// Merge Two Sorted Lists
// leetcode: https://leetcode-cn.com/problems/merge-two-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = MergeTwoSortedLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = MergeTwoSortedLists(l2.Next, l1)
		return l2
	}
}