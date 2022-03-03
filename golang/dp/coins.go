package dp

// Coins
// https://atcoder.jp/contests/dp/tasks/dp_i
func Coins(N int, ps []float64) float64 {
	dp := make([][]float64, N+1)
	for i := range dp {
		dp[i] = make([]float64, N+1)
	}

	dp[0][0] = 1
	for i := 1; i <= N; i++ {
		p := ps[i-1]
		for j := 0; j <= i; j++ {
			dp[i][j] = dp[i-1][j] * (1-p)
			if j > 0 {
				dp[i][j] += dp[i-1][j-1] * p
			}
		}
	}

	ans := 0.0
	for i := 0; i <= N; i++ {
		if N - i < i {
			ans += dp[N][i]
		}
	}
	return ans
}