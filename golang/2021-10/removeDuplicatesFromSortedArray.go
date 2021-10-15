package _021_10

// Remove Duplicates from Sorted Array
// leetcode: https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicatesFromSortedArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	preVal, preIdx := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == preVal {
			continue
		} else {
			nums[preIdx] = nums[i]
			preVal = nums[i]
			preIdx++
		}
	}
	return preIdx
}

func RemoveDuplicateFromSortedArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < l; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}