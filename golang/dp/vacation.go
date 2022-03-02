package dp

// Vacation
// https://atcoder.jp/contests/dp/tasks/dp_c
func Vacation(N int, happiness [][]int) int {
	n := len(happiness[0])
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, n)
		if i == 0 {
			dp[i] = happiness[0]
		}
	}

	for i := 1; i < N; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if k != j {
					dp[i][j] = max(dp[i-1][k]+happiness[i][j], dp[i][j])
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(dp[N-1][i], ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}