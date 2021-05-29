package _021_05

import (
	"math"
)

// Max Sum of Rectangle No Larger Than K
// leetcode: https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/
// Given an m x n matrix matrix and an integer k, return the max sum of a rectangle in the matrix such that its sum is no larger than k.
// It is guaranteed that there will be a rectangle with a sum no larger than k.

func MaxSumOfRectangleNoLargerThanK(matrix [][]int, k int) int {
	maxSum := math.MinInt64
	for i := range matrix {
		prefixSum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] {
			for j, n := range row {
				prefixSum[j] += n
			}
			maxSum = max(maxSum, findCurListMaxSum(prefixSum, k))
			if maxSum == k {
				return k
			}
		}
	}
	return maxSum
}

func findCurListMaxSum(nums []int, k int) int {
	curMax := math.MinInt64
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		for j := 0; j < i; j++ {
			temp := preSum[i] - preSum[j]
			if temp <= k {
				curMax = max(curMax, temp)
			}
			if curMax == k {
				return curMax
			}
		}
	}
	return curMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}