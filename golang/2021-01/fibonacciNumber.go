package _021_01

// Fibonacci Number
// leetCode: https://leetcode-cn.com/problems/fibonacci-number/
// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1.

// 递归
func FibonacciNumber(n int) int {
	if n < 2 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

// dp
func FibonacciNumber2(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}