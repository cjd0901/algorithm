package _022_06

import "sort"

// 摆动排序II
// leetcode: https://leetcode.cn/problems/wiggle-sort-ii/
func WiggleSortII(nums []int) {
	l := len(nums)
	temp := append([]int(nil), nums...)
	sort.Ints(temp)
	sub := (l + 1) / 2
	for i, j, k := 0, sub-1, l-1; i < l; i += 2 {
		nums[i] = temp[j]
		if i+1 < l {
			nums[i+1] = temp[k]
		}
		j--
		k--
	}
}
