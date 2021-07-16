package _021_07

// Find the Number in the Sorted Array
// leetcode:https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func FindTheNumberInTheSortedArray(nums []int, target int) int {
	count := 0
	length := len(nums)
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + (r-l) / 2
		n := nums[mid]
		if n == target {
			if mid + 1 < length {
				for i := mid + 1; i < length; i++ {
					if nums[i] != target {
						break
					}
					count++
				}
			}
			if mid >= 0 {
				for i := mid; i >= 0; i-- {
					if nums[i] != target {
						break
					}
					count++
				}
			}
			return count
		} else if n > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return count
}