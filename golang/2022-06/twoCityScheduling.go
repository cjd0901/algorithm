package _022_06

import (
	"sort"
)

// 两地调度
// leetcode: https://leetcode.cn/problems/two-city-scheduling/
func TwoCityScheduling(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[j][0] - costs[j][1]) > (costs[i][0] - costs[i][1])
	})
	l := len(costs)
	ans := 0
	for i := 0; i < l; i++ {
		if i < l/2 {
			ans += costs[i][0]
		} else {
			ans += costs[i][1]
		}
	}
	return ans
}
