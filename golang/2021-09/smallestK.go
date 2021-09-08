package _021_09

import (
	"container/heap"
)

// Smallest K
// leetcode: https://leetcode-cn.com/problems/smallest-k-lcci/
func SmallestK(arr []int, k int) []int {
	if k == 0 {
		return arr
	}
	h := hp(arr[:k])
	heap.Init(&h)
	for _, n := range arr[k:] {
		if h[0] > n {
			h[0] = n
			heap.Fix(&h, 0)
		}
	}
	return h
}

type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *hp) Pop() interface{} {
	a := *h
	temp := a[len(a)-1]
	*h = a[:len(a)-1]
	return temp
}