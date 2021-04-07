package _021_04

// Search in Rotated Sorted Array II
// leetCode: https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
// Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums

// 一次遍历
//func SearchInRotatedSortedArray(nums []int, target int) bool {
//	for i := 0; i < len(nums); i ++ {
//		if target == nums[i] {
//			return true
//		}
//	}
//	return false
//}

// 二分查找
func SearchInRotatedSortedArray(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if nums[start] == nums[mid] {
			start ++
			continue
		}
		if nums[start] < nums[mid] {
			if nums[mid] > target && nums[start] <= target {
				end = mid - 1
			}else {
				start = mid + 1
			}
		}else {
			if nums[mid] < target && nums[end] >= target {
				start = mid + 1
			}else {
				end = mid - 1
			}
		}
	}
	return false
}