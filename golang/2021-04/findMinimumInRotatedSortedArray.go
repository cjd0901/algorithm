package _021_04

// Find Minimum in Rotated Sorted Array
// leetCode: https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.


func FindMinimumInRotatedSortedArray(nums []int) int {
	low, high := 0, len(nums) - 1
	for low < high {
		pivot := low + (high - low) / 2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}

//func FindMinimumInRotatedSortedArray(nums []int) int {
//	ans := math.MaxInt8
//	for _, n := range nums {
//		if n < ans {
//			ans = n
//		}
//	}
//	return ans
//}