package golang

import (
	"strconv"
)

// Monotone Increasing Digits
// leetCode: https://leetcode-cn.com/problems/monotone-increasing-digits/
// Given a non-negative integer N, find the largest number that is less than or equal to N with monotone increasing digits.
// (Recall that an integer has monotone increasing digits if and only if each pair of adjacent digits x and y satisfy x <= y.)

func MonotoneIncreasingDigits(N int) (I int) {
	I = N
	str := strconv.Itoa(N)
	var getMax func(s string)
	getMax = func(s string) {
		for i := range s {
			if i == len(s) - 1 {
				return
			}
			if s[i] > s[i+1] {
				subStr, _ := strconv.Atoi(s[i+1:])
				newN, _ := strconv.Atoi(s)
				I = newN - (subStr+1)
				getMax(strconv.Itoa(I))
			}
		}
	}
	getMax(str)
	return
}

// 官方题解
func OfficialMonotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
			//fmt.Println(s, i)
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
			//fmt.Println(s, i)
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}
