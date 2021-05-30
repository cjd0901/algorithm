package _021_05

// Linked List Components
// leetcode: https://leetcode-cn.com/problems/linked-list-components/
// We are given head, the head node of a linked list containing unique integer values.
// We are also given the list nums, a subset of the values in the linked list.
// Return the number of connected components in nums, where two values are connected if they appear consecutively in the linked list.

type ListNode struct {
	Val int
	Next *ListNode
}

func LinkedListComponents(head *ListNode, nums []int) int {
	numMap := make(map[int]bool)
	for _, n := range nums {
		numMap[n] = true
	}
	pre, ans := false, 0
	for head != nil {
		if numMap[head.Val] && !pre {
			ans++
			pre = true
		} else if !numMap[head.Val] {
			pre = false
		}
		head = head.Next
	}
	return ans
}