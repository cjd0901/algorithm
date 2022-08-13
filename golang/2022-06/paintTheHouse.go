package _022_06

import "math"

// 粉刷房子
// leetcode: https://leetcode.cn/problems/JEj789/
func PaintTheHouse(costs [][]int) int {
	m := len(costs)
	n := len(costs[0])
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			c := costs[i][j]
			costs[i][j] = math.MaxInt
			for k := 0; k < n; k++ {
				if k != j {
					costs[i][j] = min(costs[i-1][k]+c, costs[i][j])
				}
			}
		}
	}

	ans := math.MaxInt
	for _, c := range costs[m-1] {
		if c < ans {
			ans = c
		}
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
