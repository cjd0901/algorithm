package _021_04

import "math"

// Remove Duplicates from Sorted Array II
// leetCode: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
// Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.
// Do not allocate extra space for another array; you must do this by modifying the input array in-place with O(1) extra memory.

func RemoveDuplicates(nums []int) int {
	prev, cur, curL := 0, math.MinInt16, 0
	for _, v := range nums {
		if cur != v {
			curL = 1
		}
		if curL <= 2 {
			nums[prev] = v
			prev ++
		}
		curL ++
		cur = v
	}
	return prev
}