package _021_04

// Swap Nodes in Pairs
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// Given a linked list, swap every two adjacent nodes and return its head.

type ListNode struct {
	 Val int
	 Next *ListNode
}

func SwapNodesInPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = SwapNodesInPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func SwapNodesInPairs2(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}