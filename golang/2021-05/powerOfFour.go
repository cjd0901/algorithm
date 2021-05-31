package _021_05

// Power of Four
// leetcode: https://leetcode-cn.com/problems/power-of-four/
// Given an integer n, return true if it is a power of four. Otherwise, return false.
// An integer n is a power of four, if there exists an integer x such that n == 4x.

func PowerOfFour(n int) bool {
	return n > 0 && n & (n-1) == 0 && n % 3 == 1
}