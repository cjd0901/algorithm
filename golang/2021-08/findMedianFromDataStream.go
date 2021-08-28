package _021_08

import (
	"container/heap"
	"sort"
)

// Find Median from Data Stream
// leetcode: https://leetcode-cn.com/problems/find-median-from-data-stream/
type MedianFinder struct {
	minQ, maxQ hp
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}


func (this *MedianFinder) AddNum(num int)  {
	minQ, maxQ := &this.minQ, &this.maxQ
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if minQ.Len() < maxQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}


func (this *MedianFinder) FindMedian() float64 {
	minQ, maxQ := this.minQ, this.maxQ
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0])/2
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}