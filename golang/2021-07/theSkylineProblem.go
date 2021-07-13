package _021_07

import (
	"container/heap"
	"sort"
)

// the Skyline Problem
// leetcode:https://leetcode-cn.com/problems/the-skyline-problem/

type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func (h hp) find(n int) int {
	for i, height := range h {
		if height == n {
			return i
		}
	}
	return len(h)
}

func TheSkylineProblem(buildings [][]int) [][]int {
	var ans [][]int
	points := make([][]int, 0, len(buildings)*2)
	for _, building := range buildings {
		points = append(points, []int{building[0], -building[2]})
		points = append(points, []int{building[1], building[2]})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})

	hp := hp{}
	prev := 0
	heap.Push(&hp, prev)
	for _, point := range points {
		p, h := point[0], point[1]
		if h < 0 {
			heap.Push(&hp, -h)
		} else {
			idx := hp.find(h)
			heap.Remove(&hp, idx)
		}

		cur := hp[0]
		if cur != prev {
			ans = append(ans, []int{p, cur})
			prev = cur
		}
	}

	return ans
}