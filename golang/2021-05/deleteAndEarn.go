package _021_05

// Delete And Earn
// leetcode: https://leetcode-cn.com/problems/delete-and-earn/
// Given an array nums of integers, you can perform operations on the array.
// In each operation, you pick any nums[i] and delete it to earn nums[i] points.
// After, you must delete every element equal to nums[i] - 1 or nums[i] + 1.
// You start with 0 points.
// Return the maximum number of points you can earn by applying such operations.

// dp:数组
func DeleteAndEarn(nums []int) int {
	maxVal := 0
	for _, num := range nums {
		maxVal = max(num, maxVal)
	}
	sums := make([]int, maxVal+1)
	for _, num := range nums {
		sums[num] += num
	}
	dp := make([]int, maxVal+1)
	dp[0], dp[1] = sums[0], max(sums[0], sums[1])
	for i := 2; i < len(sums); i ++ {
		dp[i] = max(dp[i-2]+sums[i], dp[i-1])
	}
	return dp[maxVal]
}

// dp:滚动数组
//func DeleteAndEarn(nums []int) int {
//	maxVal := 0
//	for _, num := range nums {
//		maxVal = max(num, maxVal)
//	}
//	sums := make([]int, maxVal+1)
//	for _, num := range nums {
//		sums[num] += num
//	}
//	first, second := sums[0], max(sums[0], sums[1])
//	for i := 2; i < len(sums); i ++ {
//		first, second = second, max(first+sums[i], second)
//	}
//	return second
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}