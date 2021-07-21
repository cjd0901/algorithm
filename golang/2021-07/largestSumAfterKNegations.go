package _021_07

import (
	"sort"
)

// Maximize Sum Of Array After K Negations
// leetcode: https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/
func LargestSumAfterKNegations(nums []int, k int) int {
	sum := 0
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	if k > 0 && k % 2 == 1 {
		nums[len(nums) - 1] *= -1
	}
	for _, n := range nums {
		sum += n
	}
	return sum
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}