package _021_08

// Remove Zero Sum Consecutive Nodes from Linked List
// leetcode: https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveZeroSumSublists(head *ListNode) *ListNode {
	temp := &ListNode{0, head}
	m := make(map[int]*ListNode)
	sum := 0
	for p := temp; p != nil; p = p.Next {
		sum += p.Val
		m[sum] = p
	}
	sum = 0
	for p := temp; p != nil; p = p.Next {
		sum += p.Val
		p.Next = m[sum].Next
	}
	return temp.Next
}