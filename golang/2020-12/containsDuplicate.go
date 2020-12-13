package golang

// Contains Duplicate
// leetCode: https://leetcode-cn.com/problems/contains-duplicate/

// Given an array of integers, find if the array contains any duplicates.
// Your function should return true if any value appears at least twice in the array,
// and it should return false if every element is distinct.

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] ++
		if m[n] > 1 {
			return true
		}
	}
	return false
}