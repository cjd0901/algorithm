package _021_05

import "sort"

// 3Sum
// leetcode: https://leetcode-cn.com/problems/3sum/
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
			if sum > 0 {
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
	return ans
}