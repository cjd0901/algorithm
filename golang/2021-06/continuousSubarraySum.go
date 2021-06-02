package _021_06

// Continuous Subarray Sum
// leetcode:https://leetcode-cn.com/problems/continuous-subarray-sum/
// Given an integer array nums and an integer k, return true if nums has a continuous subarray of size at least two whose elements sum up to a multiple of k, or false otherwise.
// An integer x is a multiple of k if there exists an integer n such that x = n * k.
// 0 is always a multiple of k.

func ContinuousSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	remainderMap := map[int]int{0: -1}
	remainder := 0
	for i, num := range nums {
		remainder = (remainder + num) % k
		if index, ok := remainderMap[remainder]; ok {
			if i - index >= 2 {
				return true
			}
		} else {
			remainderMap[remainder] = i
		}
	}
	return false
}

// 前缀和 + 数组 超时
func ContinuousSubarraySum2(nums []int, k int) bool {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1]+nums[i-1]
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			sum := prefixSum[i] - prefixSum[j]
			if sum % k == 0 && i-j >= 2 {
				return true
			}
		}
	}
	return false
}