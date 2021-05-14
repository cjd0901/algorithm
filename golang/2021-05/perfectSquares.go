package _021_05

import (
	"math"
)

// Perfect Squares
// leetcode: https://leetcode-cn.com/problems/perfect-squares/
// Given an integer n, return the least number of perfect square numbers that sum to n.

func PerfectSquares(n int) int {
	squareNums := make([]int, int(math.Sqrt(float64(n)))+1)
	for i := 1; i < len(squareNums); i++ {
		squareNums[i] = i*i
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < len(squareNums) && i - squareNums[j] >= 0; j++ {
			dp[i] = min(dp[i], dp[i-squareNums[j]]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}