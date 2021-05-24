package _021_05

import (
	"math"
)

// Strange Printer
// leetcode: https://leetcode-cn.com/problems/strange-printer/
// There is a strange printer with the following two special properties:
// The printer can only print a sequence of the same character each time.
// At each turn, the printer can print new characters starting from and ending at any place and will cover the original existing characters.
// Given a string s, return the minimum number of turns the printer needed to print it.

func StrangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n-1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i+1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			}else {
				dp[i][j] = math.MaxInt64
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}