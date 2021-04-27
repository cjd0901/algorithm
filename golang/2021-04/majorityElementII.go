package _021_04

// Majority Element II
// leetCode: https://leetcode-cn.com/problems/majority-element-ii/
// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

// TODO: 摩尔投票法
func MajorityElementII(nums []int) []int {
	n := len(nums) / 3
	ans := []int{}
	m := make(map[int]int)
	for _, num := range nums {
		m[num] ++
	}
	for k, v := range m {
		if v > n {
			ans = append(ans, k)
		}
	}
	return ans
}