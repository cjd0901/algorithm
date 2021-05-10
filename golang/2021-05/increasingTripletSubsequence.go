package _021_05

import (
	"math"
)

// Increasing Triplet Subsequence
// leetcode: https://leetcode-cn.com/problems/increasing-triplet-subsequence/
// Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func IncreasingTripletSubsequence(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	small, mid := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= small {
			small = num
		}else if num <= mid {
			mid = num
		}else if num > mid {
			return true
		}
	}
	return false
}