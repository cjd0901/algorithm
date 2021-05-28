package _021_05

// Total Hamming Distance
// leetcode: https://leetcode-cn.com/problems/total-hamming-distance/
// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
// Given an integer array nums, return the sum of Hamming distances between all the pairs of the integers in nums.

func TotalHammingDistance(nums []int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < 30; i++ {
		c := 0
		for _, v := range nums {
			c += v >> i & 1
		}
		ans += c * (n - c)
	}
	return ans
}