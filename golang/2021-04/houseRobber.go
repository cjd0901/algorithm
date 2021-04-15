package _021_04

import "fmt"

// House Robber
// leetCode: https://leetcode-cn.com/problems/house-robber/
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

func HouseRobber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}
	fmt.Println(dp)
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}