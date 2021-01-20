package _021_01

import (
	"math"
	"sort"
)

// Maximum Product of Three Numbers
// leetCode: https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

func MaximumProduct(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	return int(math.Max(float64(nums[0] * nums[1] * nums[n-1]), float64(nums[n-1] * nums[n-2] * nums[n-3])))
}