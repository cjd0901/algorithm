package _021_07

// Reorder List
// leetcode: https://leetcode-cn.com/problems/reorder-list/
type ListNode struct {
	Val int
	Next *ListNode
}
func reorderList(head *ListNode) {
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	l, r := 0, len(list)-1
	for l < r {
		list[l].Next = list[r]
		l++
		if l == r {
			break
		}
		list[r].Next = list[l]
		r--
	}
	list[l].Next = nil
}