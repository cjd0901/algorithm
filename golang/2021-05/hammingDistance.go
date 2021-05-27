package _021_05

// Hamming Distance
// leetcode:https://leetcode-cn.com/problems/hamming-distance/
// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
// Given two integers x and y, return the Hamming distance between them.

func HammingDistance(x, y int) int {
	ans := 0
	for s := x^y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return ans
}