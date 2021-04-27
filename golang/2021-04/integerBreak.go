package _021_04

// Integer Break
// leetCode: https://leetcode-cn.com/problems/integer-break/
// Given an integer n, break it into the sum of k positive integers, where k >= 2, and maximize the product of those integers.
// Return the maximum product you can get.

func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j * (i-j), j * dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}