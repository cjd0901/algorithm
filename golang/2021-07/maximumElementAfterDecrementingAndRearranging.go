package _021_07

import (
	"sort"
)

// Maximum Element After Decreasing and Rearranging
// leetcode:https://leetcode-cn.com/problems/maximum-element-after-decreasing-and-rearranging/
func MaximumElementAfterDecrementingAndRearranging(arr []int) int {
	max, n := 0, len(arr)
	if n == 1 {
		return 1
	}
	sort.Ints(arr)
	if arr[0] != 1 {
		arr[0] = 1
	}
	for i := 1; i < n; i++ {
		if arr[i] - arr[i-1] > 1 {
			max = arr[i-1] + 1
			arr[i] = arr[i-1] + 1
		} else {
			max = arr[i]
		}
	}
	return max
}