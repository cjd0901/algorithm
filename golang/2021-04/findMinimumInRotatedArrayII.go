package _021_04

import "math"

// Find Minimum in Rotated Sorted Array II
// leetCode: https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
// Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.

func FindMinimumInRotatedArrayII(nums []int) int {
	ans := math.MaxInt8
	for _, n := range nums {
		if n < ans {
			ans = n
		}
	}
	return ans
}