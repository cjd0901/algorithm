package _021_04

// Minimum Path Sum
// leetCode: https://leetcode-cn.com/problems/minimum-path-sum/
// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
// Note: You can only move either down or right at any point in time.

func MinimumPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			for j := 0; j < n; j++ {
				if j == 0 {
					dp[i][j] = grid[i][j]
				}else {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				}
			}
		} else if i > 0 {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j] + grid[i][j], dp[i][j-1] + grid[i][j])
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}