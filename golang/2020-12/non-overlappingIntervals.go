package golang

import "sort"

// Non-overlapping Intervals
// leetCode: https://leetcode-cn.com/problems/non-overlapping-intervals/
// Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

func EraseOverLapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}