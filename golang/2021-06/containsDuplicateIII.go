package _021_06

// Contains Duplicate III
// leetcode: https://leetcode-cn.com/problems/contains-duplicate-iii/
// Given an integer array nums and two integers k and t, return true if there are two distinct indices i and j in the array such that abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.

func ContainsDuplicateIII(nums []int, k, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j <= min(i+k, len(nums)-1); j++ {
			if abs(nums[j] - nums[i]) <= t {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}