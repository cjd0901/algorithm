package _021_03

import (
	"sort"
)

// Subsets II
// leetCode: https://leetcode-cn.com/problems/subsets-ii/
// Given an integer array nums that may contain duplicates, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets. Return the solution in any order.

func SubsetsII(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	t := []int{}
	n := len(nums)
	used := make([]bool, n)
	var dfs func(index int)
	dfs = func(index int) {
		ans = append(ans, append([]int(nil), t...))
		for i := index; i < n; i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			t = append(t, nums[i])
			used[i] = true
			dfs(i+1)
			used[i] = false
			t = t[:len(t)-1]
		}
	}
	dfs(0)
	return ans
}