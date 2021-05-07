package _021_03

// Rotate List
// leetCode: https://leetcode-cn.com/problems/rotate-list/
// Given the head of a linked list, rotate the list to the right by k places.

type ListNode struct {
	 Val int
	 Next *ListNode
}

// 官方答案：形成闭环 跑的时候 内存溢出
func RotateList(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}

	move := n - k%n
	if move == n {
		return head
	}

	iter.Next = head
	for move > 0 {
		iter = iter.Next
		move--
	}
	ans := iter.Next
	iter.Next = nil
	return ans
}