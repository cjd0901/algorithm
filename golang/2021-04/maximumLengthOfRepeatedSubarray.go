package _021_04

// Maximum Length of Repeated Subarray
// leetCode: https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
// Given two integer arrays nums1 and nums2, return the maximum length of a subarray that appears in both arrays.

func MaximumLengthOfRepeatedSubarray(nums1, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}else {
				dp[i][j] = 0
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}