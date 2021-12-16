package _021_12

// Odd Even Linked List
// leetcode: https://leetcode-cn.com/problems/odd-even-linked-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	o, e := head, head.Next
	i := 1
	head = head.Next.Next
	for head != nil {
		if i % 2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
		i++
	}
	even.Next = nil
	odd.Next = e
	return o
}
