package _021_08

import (
	"container/heap"
	"sort"
)

// The K Weakest Rows in a Matrix
// leetcode: https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
func TheKWeakestRowsInAMatrix(mat [][]int, k int) []int {
	h := hp{}
	for i, v := range mat {
		pow := sort.Search(len(v), func(i int) bool {
			return v[i] == 0
		})
		h = append(h, pair{
			pow: pow,
			row: i,
		})
	}
	heap.Init(&h)
	ans := []int{}
	for k > 0 {
		ans = append(ans, heap.Pop(&h).(pair).row)
		k--
	}
	return ans
}

type pair struct {
	pow, row int
}

type hp []pair

func (hp hp) Len() int {return len(hp)}
func (hp hp) Less(i, j int) bool {
	a, b := hp[i], hp[j]
	return a.pow < b.pow || a.pow == b.pow && a.row < b.row
}
func (hp hp) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}
func (hp *hp) Push(v interface{}) {
	*hp = append(*hp, v.(pair))
}
func (hp *hp) Pop() interface{} {
	a := *hp
	v := a[len(a)-1]
	*hp = a[:len(a)-1]
	return v
}