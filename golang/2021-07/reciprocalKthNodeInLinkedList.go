package _021_07

// 链表中倒数第k个节点
// leetcode: https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
func reciprocalKthNodeInLinkedList(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}