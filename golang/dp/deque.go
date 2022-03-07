package dp

// Deque
// https://atcoder.jp/contests/dp/tasks/dp_l
type pair struct {
	begin, end int
}

func Deque(N int, A []int) int {
	dp := make([][]pair, N)
	for i := range dp {
		dp[i] = make([]pair, N)
	}

	for i, a := range A {
		dp[i][i] = pair{
			begin: a,
			end:   0,
		}
	}

	for i := N-2; i >= 0; i-- {
		for j := i+1; j < N; j++ {
			begin := dp[i+1][j].end + A[i]
			end := dp[i][j-1].end + A[j]
			dp[i][j].begin = max(begin, end)

			if begin > end {
				dp[i][j].end = dp[i+1][j].begin
			} else {
				dp[i][j].end = dp[i][j-1].begin
			}
		}
	}
	return dp[0][N-1].begin - dp[0][N-1].end
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}