package _021_05

import (
	"math"
	"sort"
)

// 3Sum Closest
// leetcode: https://leetcode-cn.com/problems/3sum-closest/
// Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
// Return the sum of the three integers.
// You may assume that each input would have exactly one solution.

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	best := math.MaxInt32
	update := func(s int) {
		if abs(s - target) < abs(best - target) {
			best = s
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				tempK := k - 1
				for j < tempK && nums[tempK] == nums[k] {
					tempK--
				}
				k = tempK
			}else {
				tempJ := j + 1
				for tempJ < k && nums[tempJ] == nums[j] {
					tempJ++
				}
				j = tempJ
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}