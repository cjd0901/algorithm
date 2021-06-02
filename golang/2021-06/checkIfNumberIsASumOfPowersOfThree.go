package _021_06

// Check if Number is a Sum of Powers of Three
// leetcode: https://leetcode-cn.com/problems/check-if-number-is-a-sum-of-powers-of-three/
// Given an integer n, return true if it is possible to represent n as the sum of distinct powers of three. Otherwise, return false.
// An integer y is a power of three if there exists an integer x such that y == 3x.

func CheckIfNumberIsASumOfPowersOfThree(n int) bool {
	for n > 0 {
		if n % 3 == 1 || n % 3 == 0 {
			n /= 3
		}else {
			return false
		}
	}
	return true
}

func CheckIfNumberIsASumOfPowersOfThree2(n int) bool {
	for n > 0 {
		if n % 3 == 2 {
			return false
		} else {
			n /= 3
		}
	}
	return true
}

func CheckIfNumberIsASumOfPowersOfThree3(n int) bool {
	if n == 1 {
		return true
	}
	return n % 3 != 2 && CheckIfNumberIsASumOfPowersOfThree3(n/3)
}