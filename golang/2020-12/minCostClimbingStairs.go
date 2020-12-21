package golang

// Min Cost Climbing Stairs
// leetCode: https://leetcode-cn.com/problems/min-cost-climbing-stairs/
// On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
// Once you pay the cost, you can either climb one or two steps.
// You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

// 动态规划
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
