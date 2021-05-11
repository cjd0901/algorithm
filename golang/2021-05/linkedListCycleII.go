package _021_05

// Linked List Cycle II
// leetcode: https://leetcode-cn.com/problems/linked-list-cycle-ii/
// Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
// Note that pos is not passed as a parameter.
// Notice that you should not modify the linked list.

type ListNode struct {
	Val int
	Next *ListNode
}

func LinkedListCycleII(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// hash
func LinkedListCycleII2(head *ListNode) *ListNode {
	nMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := nMap[head]; ok {
			return head
		}
		nMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}