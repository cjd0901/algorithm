package _021_04

// Ugly Number
// leetcode: https://leetcode-cn.com/problems/ugly-number/
// An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.
// Given an integer n, return true if n is an ugly number.

func UglyNumber(n int) bool {
	factors := []int{2, 3, 5}
	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n % f == 0 {
			n /= f
		}
	}
	return n==1
}
