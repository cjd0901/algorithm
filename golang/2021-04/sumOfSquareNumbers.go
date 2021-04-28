package _021_04

import (
	"math"
)

// Sum of Square Numbers
// leetCode: https://leetcode-cn.com/problems/sum-of-square-numbers/
// Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.

func SumOfSquareNumbers(c int) bool {
	for a := 0; a*a <= c; a++ {
		rt := math.Sqrt(float64(c - a*a))
		if rt == math.Floor(rt) {
			return true
		}
	}
	return false
}

func SumOfSquareNumbers2(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}