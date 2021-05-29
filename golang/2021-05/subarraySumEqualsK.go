package _021_05

// Subarray Sum Equals K
// leetcode: https://leetcode-cn.com/problems/subarray-sum-equals-k/
// Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.

func SubarraySumEqualsK(nums []int, k int) int {
	count, pre := 0, 0
	mp := make(map[int]int)
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			count += mp[pre-k]
		}
		mp[pre]++
	}
	return count
}

func SubarraySumEqualsK2(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}