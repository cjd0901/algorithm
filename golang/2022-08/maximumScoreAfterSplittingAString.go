package main

import "strings"

func MaxScore(s string) int {
	ans := 0
	l := len(s)
	for i := 0; i < l-1; i++ {
		left, right := 0, 0
		for j := 0; j <= i; j++ {
			if s[j] == '0' {
				left++
			}
		}
		for k := i + 1; k < l; k++ {
			if s[k] == '1' {
				right++
			}
		}
		ans = max(ans, left+right)
	}
	return ans
}

func MaxScore2(s string) int {
	score := int('1'-s[0]) + strings.Count(s[1:], "1")
	ans := score
	for _, c := range s[1 : len(s)-1] {
		if c == '1' {
			score--
		} else {
			score++
		}
		ans = max(ans, score)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
