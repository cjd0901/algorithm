package _021_08

import (
	"container/heap"
)

// Super Ugly Number
// leetcode: https://leetcode-cn.com/problems/super-ugly-number/
func SuperUglyNumber(n int, primes []int) int {
	h := hp{}
	for i := range primes {
		heap.Push(&h, pair{
			val: primes[i],
			i:   i,
			idx: 0,
		})
	}
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; {
		p := heap.Pop(&h).(pair)
		if p.val != ans[i-1] {
			ans[i] = p.val
			i++
		}
		heap.Push(&h, pair{
			val: primes[p.i]*ans[p.idx+1],
			i: p.i,
			idx: p.idx+1,
		})
	}
	return ans[n-1]
}

func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

type hp []pair

type pair struct {
	val, i, idx int
}