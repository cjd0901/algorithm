package _021_10

// Remove Nth Node From End of List
// leetcode: https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthNodeFromEndOfList(head *ListNode, n int) *ListNode {
	temp := &ListNode{
		Next: head,
	}
	slow, fast := temp, temp
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return temp.Next
}
