package _021_08

import (
	"math"
)

// Network Delay Time
// leetcode: https://leetcode-cn.com/problems/network-delay-time/

func NetworkDelayTime(times [][]int, n int, k int) int {
	const inf = math.MaxInt32
	adjMatrix := make([][]int, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, n)
		for j := range adjMatrix[i] {
			adjMatrix[i][j] = inf
		}
	}
	for _, t := range times {
		x, y := t[0] - 1, t[1] - 1
		adjMatrix[x][y] = t[2]
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0
	used := make([]bool, n)
	for j := 0; j < n; j++ {
		x := -1
		for i, u := range used {
			if !u && (x == -1 || dist[i] < dist[x]) {
				x = i
			}
		}
		used[x] = true
		for i, time := range adjMatrix[x] {
			dist[i] = min(dist[i], time+dist[x])
		}
	}
	ans := 0
	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}