package _021_01

import (
	"strconv"
)

// Summary Ranges
// leetCode: https://leetcode-cn.com/problems/summary-ranges/
// You are given a sorted unique integer array nums.
// Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
// That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.
// Each range [a,b] in the list should be output as:
// "a->b" if a != b
// "a" if a == b

func SummaryRanges(nums []int) []string {
	ans := make([]string, 0)
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans,s)
	}
	return ans
}
