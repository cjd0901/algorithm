package dp

import "math"

// Slimes
// https://atcoder.jp/contests/dp/tasks/dp_n
func Slimes(N int, A []int) int {
	prefixSum := make([]int, N+1)
	for i, a := range A {
		prefixSum[i+1] = prefixSum[i] + a
	}

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}

	for l := 2; l <= N; l++ {
		for i := 0; i+l <= N; i++ {
			j := i + l
			dp[i][j] = math.MaxInt64
			for k := i+1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k] + dp[k][j] + prefixSum[j] - prefixSum[i])
			}
		}
	}

	return dp[0][N]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}