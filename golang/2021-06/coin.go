package _021_06

// Coin
// leetcode:https://leetcode-cn.com/problems/coin-lcci/

func Coin(n int) int {
	coins := []int{25, 10, 5, 1}
	dp := make([]int, n+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[n] % (1e9 + 7)
}