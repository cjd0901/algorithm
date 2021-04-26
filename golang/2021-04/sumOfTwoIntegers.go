package _021_04

// Sum of Two Integers
// leetCode: https://leetcode-cn.com/problems/sum-of-two-integers/
// Given two integers a and b, return the sum of the two integers without using the operators + and -.

func SumOfTwoIntegers(a, b int) int {
	for a != 0 {
		aT := (a & b) << 1
		bT := a ^ b
		a = aT
		b = bT
	}
	return a | b
}