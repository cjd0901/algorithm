package _021_05

// Single Element in a Sorted Array
// leetcode: https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
// You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once. Find this single element that appears only once.
// Follow up: Your solution should run in O(log n) time and O(1) space.

func SingleElementInASortedArray(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := lo + (hi - lo) / 2
		if mid % 2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		}else {
			hi = mid
		}
	}
	return nums[lo]
}

func SingleElementInASortedArray3(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := lo + (hi - lo) / 2
		halvesAreEven := (hi - mid) % 2 == 0
		if nums[mid] == nums[mid+1] {
			if halvesAreEven {
				lo = mid + 2
			}else {
				hi = mid - 1
			}
		}else if nums[mid] == nums[mid-1] {
			if halvesAreEven {
				hi = mid - 2
			}else {
				lo = mid + 1
			}
		}else {
			return nums[mid]
		}
	}
	return nums[lo]
}

func SingleElementInASortedArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i += 2 {
		if i == len(nums) - 1 {
			return nums[len(nums) - 1]
		}
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return 0
}
