package _021_05

import "sort"

// Max Number of K-Sum Pairs
// leetcode: https://leetcode-cn.com/problems/max-number-of-k-sum-pairs/
// You are given an integer array nums and an integer k.
// In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.
// Return the maximum number of operations you can perform on the array.

func MaxNumberOfKSumPairs(nums []int, k int) int {
	count := 0
	sort.Ints(nums)
	left, right := 0, len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum > k {
			right--
		}else if sum < k {
			left++
		}else {
			count++
			left++
			right--
		}
	}
	return count
}

// hash
func MaxNumberOfKSumPairs2(nums []int, k int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	ans := 0
	for num, count := range counts {
		if count == 0 {
			continue
		}
		if num * 2 == k {
			ans += count / 2
			counts[num] = 0
		}
		if c, ok := counts[k-num]; ok {
			ans += min(c, count)
			counts[num] = 0
			counts[k-num] = 0
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}