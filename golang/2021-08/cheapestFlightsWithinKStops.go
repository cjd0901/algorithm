package _021_08

// Cheapest Flights Within K Stops
// leetcode: https://leetcode-cn.com/problems/cheapest-flights-within-k-stops/
func CheapestFlightsWithinStops(n int, flights [][]int, src int, dst int, k int) int {
	const INF = int(1e9 * 101 + 1)
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			dp[t][i] = min(dp[t][i], dp[t-1][j]+cost)
		}
	}
	ans := INF
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == INF {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}