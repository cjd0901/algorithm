package dp

// LCS
// https://atcoder.jp/contests/dp/tasks/dp_f
func Lcs(s, t string) int {
	ns, nt := len(s), len(t)
	dp := make([][]int, ns+1)
	for i := range dp {
		dp[i] = make([]int, nt+1)
	}

	for i := 1; i <= ns; i++ {
		for j := 1; j <= nt; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]+1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[ns][nt]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}