package _021_05

import (
	"sort"
)

// Most Profit Assigning Work
// leetcode: https://leetcode-cn.com/problems/most-profit-assigning-work/

func MostProfitAssigningWork(difficulty, profit, worker []int) int {
	d2p := make(map[int]int)
	for i, d := range difficulty {
		if _, ok := d2p[d]; ok {
			d2p[d] = max(profit[i], d2p[d])
		}else {
			d2p[d] = profit[i]
		}
	}
	sort.Ints(difficulty)
	sort.Ints(worker)
	ans, i, best := 0, 0, 0
	for j := 0; j < len(worker); j++ {
		for i < len(difficulty) && worker[j] >= difficulty[i] {
			best = max(d2p[difficulty[i]], best)
			i++
		}
		ans += best
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}