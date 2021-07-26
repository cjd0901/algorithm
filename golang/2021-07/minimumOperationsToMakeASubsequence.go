package _021_07

import (
	"sort"
)

// Minimum Operations to Make a Subsequence
// leetcode: https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/
func MinimumOperationsToMakeASubsequence(target []int, arr []int) int {
	n := len(target)
	tMap := make(map[int]int)
	for i, t := range target {
		tMap[t] = i
	}
	d := []int{}
	for _, v := range arr {
		if idx, has := tMap[v]; has {
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}