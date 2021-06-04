package _021_06

// Contiguous Array
// leetcode:https://leetcode-cn.com/problems/contiguous-array/
// Given a binary array nums, return the maximum length of a contiguous subarray with an equal number of 0 and 1.

func ContiguousArray(nums []int) int {
	mp := map[int]int{0:-1}
	count, maxLength := 0, 0
	for i, num := range nums {
		if num == 1 {
			count++
		}else {
			count--
		}
		if preIndex, ok := mp[count]; ok {
			maxLength = max(maxLength, i-preIndex)
		}else {
			mp[count] = i
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}