package _021_04

// Number of Longest Increasing Subsequence
// leetCode: https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
// Given an integer array nums, return the number of longest increasing subsequences.

func NumberOfLongestIncreasingSubsequence(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}
	maxL := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				}else if dp[j] + 1 == dp[i] {
					count[i] += count[j]
				}
			}
			maxL = max(maxL, dp[i])
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxL {
			res += count[i]
		}
	}
	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}