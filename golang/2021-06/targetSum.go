package _021_06

// Target Sum
// leetcode: https://leetcode-cn.com/problems/target-sum/
// You are given an integer array nums and an integer target.
// You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.
// For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
// Return the number of different expressions that you can build, which evaluates to target.

// dp
func TargetSum(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff % 2 == 1 {
		return 0
	}
	n, neg := len(nums), diff / 2
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, neg+1)
	}
	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {
				dp[i+1][j] += dp[i][j-num]
			}
		}
	}
	return dp[n][neg]
}

// 回溯
func TargetSum2(nums []int, target int) int {
	count := 0
	var dfs func(sum, i, optNum int)
	dfs = func(sum, i, optNum int) {
		sum += optNum
		if i == len(nums) - 1 {
			if sum == target {
				count++
			}
			return
		}
		dfs(sum, i+1, -nums[i+1])
		dfs(sum, i+1, nums[i+1])
	}
	dfs(0, 0, nums[0])
	dfs(0, 0, -nums[0])
	return count
}