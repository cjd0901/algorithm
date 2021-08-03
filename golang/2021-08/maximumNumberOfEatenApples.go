package _021_08

import (
	"container/heap"
)

// Maximum Number of Eaten Apples
// leetcode: https://leetcode-cn.com/problems/maximum-number-of-eaten-apples/
func MaximumNumberOfEatenApples(apples []int, days []int) int {
	h := hp{}
	ans, i := 0, 0
	heap.Init(&h)
	for {
		if i < len(apples) {
			heap.Push(&h, apple{
				num: apples[i],
				day: i+days[i],
			})
		}
		var a apple
		for h.Len() > 0 && (a.day <= i || a.num < 1) {
			a = heap.Pop(&h).(apple)
		}
		if a.num > 0 && a.day > i {
			a.num --
			ans ++
			if a.num > 0 {
				heap.Push(&h, a)
			}
		}
		if h.Len() == 0 && i >= len(apples) {
			break
		}
		i++
	}
	return ans
}

type apple struct {
	num, day int
}


type hp []apple

func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.day < b.day
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(apple))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
