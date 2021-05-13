package _021_05

// Number of Ways to Stay in the Same Place After Some Steps
// leetcode: https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/
// You have a pointer at index 0 in an array of size arrLen.
// At each step, you can move 1 position to the left, 1 position to the right in the array or stay in the same place  (The pointer should not be placed outside the array at any time).
// Given two integers steps and arrLen, return the number of ways such that your pointer still at index 0 after exactly steps steps.
// Since the answer may be too large, return it modulo 10^9 + 7.

func NumberOfWaysToStayInTheSamePlaceAfterSomeSteps(steps, arrLen int) int {
	furthest := min(steps, arrLen-1)
	const mod = 1e9 + 7
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, furthest+1)
	}
	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j <= furthest; j++ {
			dp[i][j] = dp[i-1][j]
			if j - 1 >= 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j + 1 <= furthest {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[steps][0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}