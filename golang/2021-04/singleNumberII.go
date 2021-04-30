package _021_04

// Single Number II
// leetcode: https://leetcode-cn.com/problems/single-number-ii/
// Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

func SingleNumberII(nums []int) int {
	nMap := make(map[int]int)
	for _, n := range nums {
		nMap[n] ++
	}
	for k, v := range nMap {
		if v == 1 {
			return k
		}
	}
	return 0
}