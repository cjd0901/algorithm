package _022_06

import (
	"math/rand"
	"time"
)

// 找出第K小的数对距离(oom，不能使用额外数据结构)
// leetcode: https://leetcode.cn/problems/find-k-th-smallest-pair-distance/
func SmallestDistancePair(nums []int, k int) int {
	l := len(nums)
	pairs := make([]int, 0)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			pairs = append(pairs, abs(nums[j]-nums[i]))
		}
	}
	rand.Seed(time.Now().UnixNano())
	return quickSelect(pairs, 0, len(pairs)-1, k-1)
}

func quickSelect(pairs []int, l, r, k int) int {
	q := randPartition(pairs, l, r)
	if q == k {
		return pairs[q]
	} else if q < k {
		return quickSelect(pairs, q+1, r, k)
	}
	return quickSelect(pairs, 0, q-1, k)
}

func randPartition(pairs []int, l, r int) int {
	ra := rand.Int()%(r-l+1) + l
	pairs[l], pairs[ra] = pairs[ra], pairs[l]
	return partition(pairs, l, r)
}

func partition(pairs []int, l, r int) int {
	v := pairs[l]
	i := l
	for j := r; j > i; {
		if pairs[j] < v {
			i++
			pairs[i], pairs[j] = pairs[j], pairs[i]
		} else {
			j--
		}
	}
	pairs[l], pairs[i] = pairs[i], pairs[l]
	return i
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
