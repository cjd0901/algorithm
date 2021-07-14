package _021_07

import (
	"sort"
)

// Minimum Absolute Sum Difference
// leetcode:https://leetcode-cn.com/problems/minimum-absolute-sum-difference/

func MinimumAbsoluteSumDifference(nums1 []int, nums2 []int) int {
	n := len(nums1)
	mod := int(1e9 + 7)
	sum, maxn := 0, 0
	sorted := append([]int{}, nums1...)
	sort.Ints(sorted)
	for i, n2 := range nums2 {
		n1 := nums1[i]
		if n2 == n1 {
			continue
		}
		x := abs(n1 - n2)
		sum += x
		l, r := 0, n-1
		for l < r {
			mid := (l + r) / 2
			if sorted[mid] < n2 {
				l = mid + 1
			}else {
				r = mid
			}
		}
		temp := abs(sorted[l] - n2)
		if l > 0 {
			temp = min(temp, abs(sorted[l-1] - n2))
		}
		if temp < x {
			maxn = max(maxn, x - temp)
		}
	}
	return (sum-maxn) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}