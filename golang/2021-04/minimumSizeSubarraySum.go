package _021_04

import "math"

// Minimum Size Subarray Sum
// leetCode: https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target.
// If there is no such subarray, return 0 instead.

func MinimumSizeSubarraySum(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end - start + 1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}