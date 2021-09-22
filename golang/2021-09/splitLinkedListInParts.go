package _021_09

// Split Linked List in Parts
// leetcode: https://leetcode-cn.com/problems/split-linked-list-in-parts/
type ListNode struct {
	Val  int
	Next *ListNode
}

func SplitLinkedListInParts(root *ListNode, k int) []*ListNode {
	l := 0
	for n := root; n != nil; n = n.Next {
		l++
	}
	quotient, remainder := l/k, l%k
	ans := make([]*ListNode, k)
	for i, node := 0, root; i < k && node != nil; i++ {
		ans[i] = node
		s := quotient
		if i < remainder {
			s++
		}
		for j := 1; j < s; j++ {
			node = node.Next
		}
		node, node.Next = node.Next, nil
	}
	return ans
}
