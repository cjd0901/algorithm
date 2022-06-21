package _022_06

import (
	"unicode"
)

// 字符串中不同整数的数目
// leetcode: https://leetcode.cn/problems/number-of-different-integers-in-a-string/

func NumberOfDifferentIntegersInAString(word string) int {
	nMap := make(map[string]struct{})
	words := []rune(word)
	temp := make([]rune, 0)
	for i, u := range words {
		if unicode.IsDigit(u) &&
			(u != '0' || (len(temp) > 0 && u == '0') || (len(temp) == 0 && (i == len(words)-1 || (i < len(words)-1 && unicode.IsLetter(words[i+1]))))) {
			temp = append(temp, u)
		}
		if len(temp) > 0 && (unicode.IsLetter(u) || i == len(words)-1) {
			nMap[string(temp)] = struct{}{}
			temp = []rune{}
		}
	}
	return len(nMap)
}
