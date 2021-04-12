package _021_04

import (
	"sort"
	"strconv"
	"strings"
)

// Largest Number
// leetCode: https://leetcode-cn.com/problems/largest-number/
// Given a list of non-negative integers nums, arrange them such that they form the largest number.

func LargestNumber(nums []int) string {
	sNums := make([]string, len(nums))
	for i, num := range nums {
		sNums[i] = strconv.Itoa(num)
	}
	sort.Slice(sNums, func(i, j int) bool {
		return sNums[i] + sNums[j] >= sNums[j] + sNums[i]
	})
	ans := strings.Join(sNums, "")
	if ans[0] == '0' {
		return "0"
	}
	return ans
}