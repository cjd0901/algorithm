package _021_06

// Coin CHange 2
// leetcode: https://leetcode-cn.com/problems/coin-change-2/
// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.
// Return the number of combinations that make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return 0.

func CoinChange2(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}