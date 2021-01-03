package _021_01

// Partition List
// leetCode: https://leetcode-cn.com/problems/partition-list/
// Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.

type ListNode struct {
	Val int
	Next *ListNode
}

func PartitionList(head *ListNode, x int) *ListNode {
	small, big := &ListNode{}, &ListNode{}
	smallHead := small
	bigHead := big
	for head != nil{
		v := head.Val
		if v >= x {
			big.Next = head
			big = big.Next
		}else {
			small.Next = head
			small = small.Next
		}
		head = head.Next
	}
	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}