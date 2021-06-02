package _021_06

// Reverse Linked List II
// leetcode:https://leetcode-cn.com/problems/reverse-linked-list-ii/
// Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func ReverseLinkedListII2(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode := pre.Next
	curr := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverseLinkedList(leftNode)

	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

// 栈解法
func ReverseLinkedListII(head *ListNode, left int, right int) *ListNode {
	stack := []*ListNode{}
	i := 1
	dummyNode := &ListNode{Val: -1, Next: head}
	var pre, succ *ListNode
	for head != nil {
		if i >= left && i <= right {
			stack = append(stack, head)
		}
		if left != 1 && i == left-1 {
			pre = head
		}
		if i == right && head.Next != nil {
			succ = head.Next
		}
		head = head.Next
		i++
	}

	j := 0
	for len(stack) > 0 {
		tn := stack[len(stack)-1]
		tn.Next = nil
		stack = stack[:len(stack)-1]
		if pre == nil && j == 0 {
			dummyNode.Next = tn
		} else {
			pre.Next = tn
		}
		if succ != nil && len(stack) == 0 {
			tn.Next = succ
		}
		pre = tn
		j++
	}
	return dummyNode.Next
}