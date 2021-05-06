package _021_05

// Ones And Zeroes
// leetcode: https://leetcode-cn.com/problems/ones-and-zeroes/
// You are given an array of binary strings strs and two integers m and n.
// Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.
// A set x is a subset of a set y if all elements of x are also elements of y.

func OnesAndZeros(strs []string, m, n int) int {
	l := len(strs)
	dp := make([][][]int, l+1)
	for i := 0; i < l+1; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 1; i <= l; i++ {
		counts := count(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				zeroes := counts[0]
				ones := counts[1]
				if j >= zeroes && k >= ones {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-zeroes][k-ones]+1)
				}
			}
		}
	}
	return dp[l][m][n]
}

func count(str string) []int {
	counts := make([]int, 2)
	for _, ch := range str {
		counts[ch - '0'] ++
	}
	return counts
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}