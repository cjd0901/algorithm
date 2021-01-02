package _021_01

import (
	"fmt"
	"math"
)

// Sliding Window Maximum
// leetCode: https://leetcode-cn.com/problems/sliding-window-maximum/
// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right.
// You can only see the k numbers in the window.
// Each time the sliding window moves right by one position.
// Return the max sliding window.

func max(list []int) (int, int) {
	max := math.MinInt32
	index := len(list) - 1
	for i, num := range list{
		if num > max {
			max = num
			index = i
		}
	}
	return max, index
}

// 超出时间限制
//func MaxSlidingWindow(nums []int, k int) []int {
//	var ans []int
//	window := nums[:k]
//	ans = append(ans, max(window))
//	for i := k; i < len(nums); i++ {
//		window = window[1:]
//		window = append(window, nums[i])
//		ans = append(ans, max(window))
//	}
//	return ans
//}

// 超出时间限制
func MaxSlidingWindow2(nums []int, k int) []int {
	var ans []int
	window := nums[:k]
	m, index := max(window)
	ans = append(ans, m)
	for i := k; i < len(nums); i++ {
		window = window[1:]
		window = append(window, nums[i])
		if nums[i] >= m {
			m = nums[i]
			index = k-1
		}else {
			index--
		}
		fmt.Println(index, m)
		if index < 0 {
			m, index = max(window)
		}
		ans = append(ans, m)
	}
	return ans
}

func MaxSlidingWindow3(nums []int, k int) []int {
	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}