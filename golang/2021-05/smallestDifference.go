package _021_05

import (
	"math"
	"sort"
)

// SmallestDifference
// leetcode: https://leetcode-cn.com/problems/smallest-difference-lcci/
// Given two arrays of integers, compute the pair of values (one value in each array) with the smallest (non-negative) difference. Return the difference.

func SmallestDifference(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	aP, bP := 0, 0
	diff := math.MaxInt32
	for aP < len(a) && bP < len(b) {
		if a[aP] == b[bP] {
			return 0
		}else if a[aP] > b[bP] {
			diff = min(diff, a[aP] - b[bP])
			bP++
		}else {
			diff = min(diff, b[bP] - a[aP])
			aP++
		}
	}
	return diff
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}