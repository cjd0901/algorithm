package _021_06

// Intersection of Two Linked Lists
// leetcode: https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect.
// If the two linked lists have no intersection at all, return null.

type ListNode struct {
	Val int
	Next *ListNode
}

func IntersectionOfTwoLinkedLists(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headA
		}else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headB
		}else {
			pb = pb.Next
		}
	}
	return pa
}

// hash
func IntersectionOfTwoLinkedLists2(headA, headB *ListNode) *ListNode {
	amp := make(map[*ListNode]struct{})
	for headA != nil {
		amp[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := amp[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}