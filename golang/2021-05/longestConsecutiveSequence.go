package _021_05

import (
	"sort"
)

// Longest Consecutive Sequence
// leetcode: https://leetcode-cn.com/problems/longest-consecutive-sequence/
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

func LongestConsecutiveSequence(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			cur := num
			curL := 1
			for numSet[cur+1] {
				cur++
				curL++
			}
			longest = max(longest, curL)
		}
	}
	return longest
}

// TODO: dp 有问题
func LongestConsecutiveSequence2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n-2; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			if nums[j] == nums[j-1]+1 {
				if dp[i+1][j] != dp[i][j-1] {
					dp[i][j] = min(dp[i+1][j]+1, dp[i][j-1]+1)
				}else {
					dp[i][j] = dp[i+1][j]+1
				}
			}else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}