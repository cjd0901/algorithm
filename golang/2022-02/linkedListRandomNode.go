package _022_02

import "math/rand"

// Linked List Random Node
// leetcode: https://leetcode-cn.com/problems/linked-list-random-node/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 蓄水池抽样
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (s *Solution) GetRandom() (ans int) {
	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 { // 1/i 的概率选中（替换为答案）
			ans = node.Val
		}
		i++
	}
	return
}

//type Solution struct {
//	list []int
//}
//
//func Constructor(head *ListNode) Solution {
//	list := make([]int, 0)
//	for head != nil {
//		list = append(list, head.Val)
//		head = head.Next
//	}
//	return Solution{
//		list,
//	}
//}
//
///** Returns a random node's value. */
//func (this *Solution) GetRandom() int {
//	random := rand.Intn(len(this.list))
//	return this.list[random]
//}
