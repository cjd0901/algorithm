package _022_02

import "sort"

// Minimum Difference Between Highest and Lowest of K Scores
// leetcode: https://leetcode-cn.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
func MinimumDifference(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}
	sort.Ints(nums)
	ans := int(1e5)
	for i := 0; i <= len(nums) - k; i++ {
		diff := abs(nums[i+k-1] - nums[i])
		if diff < ans {
			ans = diff
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
