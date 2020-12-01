package golang

// 34.在排序数组中查找元素的第一个和最后一个位置
// leetCode: https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func SearchRange(nums []int, target int) []int {
	start := -1
	end := -1
	l := len(nums)
	if l == 0 {
		return []int{start, end}
	}
	for i := 0; i < l ; i++ {
		if start != -1 && end != -1 {
			break
		}
		if start == -1 && nums[i] == target {
			start = i
		}
		if end == -1 && nums[l - 1 - i] == target {
			end = l - 1 - i
		}
	}
	return []int{start, end}
}
