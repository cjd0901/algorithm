package _021_05

import (
	"math"
)

// Reverse Integer
// leetcode: https://leetcode-cn.com/problems/reverse-integer/
// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

func ReverseInteger(x int) int {
	ans := 0
	for math.Abs(float64(x)) >= 10 {
		ans = ans * 10 + x % 10
		x = x / 10
	}
	ans = ans * 10 + x % 10
	if ans >= math.MaxInt32 || ans <= math.MinInt32 {
		return 0
	}
	return ans
}