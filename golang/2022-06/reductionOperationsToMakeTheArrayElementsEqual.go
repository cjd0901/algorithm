package _022_06

import (
	"sort"
)

// 使数组元素相等的减少操作次数
// leetcode: https://leetcode.cn/problems/reduction-operations-to-make-the-array-elements-equal/
func ReductionOperationsToMakeTheArrayElementsEqual(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	ans := 0
	min := nums[0]
	for j := n - 1; j >= 0 && nums[j] != min; j-- {
		if nums[j-1] != nums[j] && nums[j] > min {
			ans += n - j
		}
	}
	return ans
}
