package _021_05

// Number of Dice Rolls With Target Sum
// leetcode: https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum/
// You have d dice and each die has f faces numbered 1, 2, ..., f.
// Return the number of possible ways (out of fd total ways) modulo 109 + 7 to roll the dice so the sum of the face-up numbers equals target.

func NumberOfDiceRollsWithTargetSum(d, f, target int) int {
	dp := make([][]int, d)
	m := min(f, target)
	mod := int(1e9 + 7)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 1; i <= m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < d; i++ {
		for j := 1; j <= target; j++ {
			for k := 0; k <= f && j-k > 0; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % mod
			}
		}
	}
	return dp[d-1][target]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}