package _021_04

// Maximum Length of Subarray With Positive Product
// https://leetcode-cn.com/problems/maximum-length-of-subarray-with-positive-product/
// Given an array of integers nums, find the maximum length of a subarray where the product of all its elements is positive.

func MaximumLengthOfSubarrayWithPositiveProduct(nums []int) int {
	n := len(nums)
	positive := make([]int, n)
	negative := make([]int, n)
	if nums[0] < 0 {
		negative[0] = 1
	}else if nums[0] > 0 {
		positive[0] = 1
	}
	maxL := positive[0]
	for i := 1; i < n; i++ {
		num := nums[i]
		if num > 0 {
			positive[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				negative[i] = negative[i-1] + 1
			}else {
				negative[i] = 0
			}
		}else if num < 0 {
			negative[i] = positive[i-1] + 1
			if negative[i-1] > 0 {
				positive[i] = negative[i-1] + 1
			}else {
				positive[i] = 0
			}
		}else {
			positive[i] = 0
			negative[i] = 0
		}
		maxL = max(maxL, positive[i])
	}
	return maxL
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}