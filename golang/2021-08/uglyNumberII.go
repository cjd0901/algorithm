package _021_08

import "container/heap"

// Ugly Number II
// leetcode: https://leetcode-cn.com/problems/ugly-number-ii/
func UglyNumberII(n int) int {
	h := hp{}
	primes := []int{2,3,5}
	for i, n := range primes {
		heap.Push(&h, pair{n, i, 0})
	}
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; {
		p := heap.Pop(&h).(pair)
		if ans[i-1] != p.val {
			ans[i] = p.val
			i++
		}
		heap.Push(&h, pair{ans[p.idx+1]*primes[p.i], p.i, p.idx+1})
	}
	return ans[n-1]
}

type hp []pair

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

type pair struct {
	val, i, idx int
}
