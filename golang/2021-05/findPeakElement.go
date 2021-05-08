package _021_05

// Find Peak Element
// leetcode: https://leetcode-cn.com/problems/find-peak-element/
// A peak element is an element that is strictly greater than its neighbors.
// Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.
// You may imagine that nums[-1] = nums[n] = -∞.

func FindPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l+r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
		}else {
			l = mid + 1
		}
	}
	return l
}

// 递归二分
func FindPeakElement3(nums []int) int {
	var binary func(l, r int) int
	binary = func(l, r int) int {
		if l == r {
			return l
		}
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			return binary(0, mid)
		}
		return binary(mid+1, r)
	}
	return binary(0, len(nums)-1)
}

// 笨比解法
func FindPeakElement2(nums []int) int {
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}