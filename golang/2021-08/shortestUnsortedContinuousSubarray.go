package _021_08

import (
	"sort"
)

// Shortest Unsorted Continuous Subarray
// leetcode:https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
func ShortestUnsortedContinuousSubarray(nums []int) int {
	sNums := append([]int(nil), nums...)
	n := len(nums)
	sort.Ints(sNums)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if nums[left] == sNums[left] {
			left++
		}
		if nums[right] == sNums[right] {
			right--
		}
	}
	if left == n && right == -1 {
		return 0
	}
	return right - left + 1
}