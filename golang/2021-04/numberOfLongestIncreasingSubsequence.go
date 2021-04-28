package _021_04

// Number of Longest Increasing Subsequence
// leetCode: https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
// Given an integer array nums, return the number of longest increasing subsequences.

// TODO: 思路错误
func NumberOfLongestIncreasingSubsequence(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++{
		if nums[i] <= nums[i-1] {
			dp[i] = dp[i-1] + 1
		}else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums) - 1]
}