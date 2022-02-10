package _022_02

import "container/heap"

// Find K Pairs with Smallest Sums
// leetcode: https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/
func FindKPairsWithSmallestSums(nums1 []int, nums2 []int, k int) [][]int {
	ans := make([][]int, 0)
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, Pair{i, 0})
	}
	for len(ans) < k && h.Len() > 0 {
		pair := heap.Pop(&h).(Pair)
		i, j := pair.i, pair.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, Pair{i, j+1})
		}
	}
	return ans
}

type Pair struct {
	i, j int
}

type hp struct {
	data []Pair
	nums1, nums2 []int
}


func (h hp) Len() int {
	return len(h.data)
}

func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}

func (h *hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *hp) Push(v interface{}) {
	h.data = append(h.data, v.(Pair))
}

func (h *hp) Pop() interface{} {
	data := h.data
	v := h.data[len(data)-1]
	h.data = h.data[:len(data)-1]
	return v
}
