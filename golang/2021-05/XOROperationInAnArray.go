package _021_05

// XOR Operation In An Array
// leetcode: https://leetcode-cn.com/problems/xor-operation-in-an-array/
// Given an integer n and an integer start.
// Define an array nums where nums[i] = start + 2*i (0-indexed) and n == nums.length.
// Return the bitwise XOR of all elements of nums.

func XOROperationInAnArray(n, start int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans = start ^ ans
		start += 2
	}
	return ans
}