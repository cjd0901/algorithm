package _022_03

// Frog 1
// https://atcoder.jp/contests/dp/tasks/dp_a
func Frog1(N int, heights []int) int {
	dp := make([]int, N)
	for i := 1; i < N; i++ {
		dp[i] = int(1e9)
	}
	for i := 1; i < N; i++ {
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+abs(heights[i] - heights[i-2]))
		}
		dp[i] = min(dp[i], dp[i-1]+abs(heights[i] - heights[i-1]))
	}
	return dp[N-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}