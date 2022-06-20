package _022_06

import (
	"math"
)

// 从数组中移除最大值和最小值
// https://leetcode.cn/problems/removing-minimum-and-maximum-from-array/

func RemovingMinimumAndMaximumFromArray(nums []int) int {
	l := len(nums)
	min, max := math.MaxInt32, math.MinInt32
	minIdx, maxIdx := 0, l-1
	for i, num := range nums {
		if num > max {
			max = num
			maxIdx = i
		}
		if num < min {
			min = num
			minIdx = i
		}
	}
	a, b := bigger(minIdx, maxIdx)
	return smaller(smaller(l-a+b+1, l-b), a+1)
}

func bigger(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}

func smaller(x, y int) int {
	if x > y {
		return y
	}
	return x
}
