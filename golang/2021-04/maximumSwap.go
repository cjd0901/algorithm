package _021_04

import (
	"strconv"
)

// Maximum Swap
// leetcode: https://leetcode-cn.com/problems/maximum-swap/
// Given a non-negative integer, you could swap two digits at most once to get the maximum valued number. Return the maximum valued number you could get.

func MaximumSwap(num int) int {
	numStr := strconv.Itoa(num)
	charArray := []byte(numStr)
	last := make([]int, 10)
	for i := 0; i < len(charArray); i++ {
		last[charArray[i] - '0'] = i
	}
	for i := 0; i < len(charArray) - 1; i++ {
		for j := 9; j > int(charArray[i] - '0'); j-- {
			if last[j] > i {
				temp := charArray[i]
				charArray[i] = charArray[last[j]]
				charArray[last[j]] = temp
				ans, _ := strconv.Atoi(string(charArray))
				return ans
			}
		}
	}
	return num
}