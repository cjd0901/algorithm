package _021_05

// Power Of Two
// leetcode:https://leetcode-cn.com/problems/power-of-two/
// Given an integer n, return true if it is a power of two. Otherwise, return false.
// An integer n is a power of two, if there exists an integer x such that n == 2x.

func PowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func PowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	for n >= 1 {
		if n == 1 {
			return true
		}
		if n % 2 != 0 {
			return false
		}
		n = n / 2.0
	}
	return false
}