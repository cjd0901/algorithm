package _021_06

// Remove Linked List Elements
// leetcode:https://leetcode-cn.com/problems/remove-linked-list-elements/
// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

type ListNode struct {
	Val int
	Next *ListNode
}

func RemoveLinkedListElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	removeElement(head, dummyNode, val)
	return dummyNode.Next
}

func removeElement(ln, pre *ListNode, val int) {
	if ln == nil {
		return
	}
	if ln.Val == val {
		pre.Next = ln.Next
		removeElement(ln.Next, pre, val)
	} else {
		removeElement(ln.Next, ln, val)
	}
}

func RemoveLinkedListElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = RemoveLinkedListElements2(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}