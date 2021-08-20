package _021_08

// Add Two Numbers
// leetcode: https://leetcode-cn.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1+n2+carry
		v := sum%10
		carry = sum/10
		if head == nil {
			head = &ListNode{
				Val: v,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: v,
			}
			tail = tail.Next
		}
	}
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}