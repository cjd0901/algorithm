package _021_01

// Rotate Array
// leetCode: https://leetcode-cn.com/problems/rotate-array/
// Given an array, rotate the array to the right by k steps, where k is non-negative.

func RotateArray(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	reverse := func(start, end int) {
		for start < end {
			temp := nums[start]
			nums[start] = nums[end]
			nums[end] = temp
			start ++
			end --
		}
	}
	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}

func RotateArray2(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	temp := nums
	temp = append(temp, nums[:n-k]...)
	temp = temp[n-k:]
	copy(nums, temp)
}