package dp

// Knapsack 1
// https://atcoder.jp/contests/dp/tasks/dp_d
func Knapsack1(N, W int, items [][]int) int {
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= items[i-1][0] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-items[i-1][0]]+items[i-1][1])
			}
		}
	}
	return dp[N][W]
}

func Knapsack12(N, W int, items [][]int) int {
	dp := make([]int, W+1)

	for i := 1; i <= N; i++ {
		for j := W; j >= items[i-1][0]; j-- {
			dp[j] = max(dp[j], dp[j-items[i-1][0]] + items[i-1][1])
		}
	}
	return dp[W]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}