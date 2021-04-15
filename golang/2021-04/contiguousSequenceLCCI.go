package _021_04


// Contiguous Sequence LCCI
// leetCode: https://leetcode-cn.com/problems/contiguous-sequence-lcci/
// You are given an array of integers (both positive and negative). Find the contiguous sequence with the largest sum. Return the sum.

func ContiguousSequenceLCCI(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//func ContiguousSequenceLCCI(nums []int) int {
//	n := len(nums)
//	if n == 1 {
//		return nums[0]
//	}
//	dp := make([]int, n)
//	dp[0] = nums[0]
//	sum := nums[0]
//	for i := 1; i < n; i++ {
//		dp[i] = max(dp[i-1]+nums[i], nums[i])
//		if dp[i] > sum {
//			sum = dp[i]
//		}
//	}
//	return sum
//}

//func ContiguousSequenceLCCI(nums []int) int {
//	n := len(nums)
//	if n == 1 {
//		return nums[0]
//	}
//	dp := make([]int, n)
//	dp[0] = nums[0]
//	sum := math.MinInt32
//	for i := 1; i < n; i++ {
//		dp[i] = max(dp[i-1]+nums[i], nums[i])
//	}
//	for _, max := range dp {
//		if max > sum {
//			sum = max
//		}
//	}
//	return sum
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}