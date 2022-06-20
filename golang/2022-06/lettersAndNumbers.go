package _022_06

import (
	"unicode"
)

// 字母与数字
// leetcode: https://leetcode.cn/problems/find-longest-subarray-lcci/
func LettersAndNumbers(array []string) []string {
	l := len(array)
	maxLength := 0
	start, end := 0, 0
	preSum := make([]int, l)
	idxMap := make(map[int]int)
	idxMap[0] = -1
	for i := 0; i < l; i++ {
		u := []rune(array[i])[0]
		if unicode.IsDigit(u) {
			if i == 0 {
				preSum[i] = -1
			} else {
				preSum[i] = preSum[i-1] - 1
			}
		} else {
			if i == 0 {
				preSum[i] = 1
			} else {
				preSum[i] = preSum[i-1] + 1
			}
		}

		if idx, ok := idxMap[preSum[i]]; ok {
			rangeLength := i - idx + 1
			if rangeLength > maxLength {
				maxLength = rangeLength
				start, end = idx+1, i
			}
		} else {
			idxMap[preSum[i]] = i
		}
	}
	if maxLength == 0 {
		return []string{}
	}
	return array[start : end+1]
}
