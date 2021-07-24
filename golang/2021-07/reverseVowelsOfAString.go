package _021_07

import (
	"unicode"
)

// Reverse Vowels of a String
// leetcode: https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
func ReverseVowelsOfAString(s string) string {
	vowelMap := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}
	runes := []rune(s)
	l, r := 0, len(runes)
	for l < r {
		if _, ok := vowelMap[unicode.ToLower(runes[l])]; ok {
			for {
				r--
				if l >= r {
					break
				}
				if _, ok := vowelMap[unicode.ToLower(runes[r])]; ok {
					runes[l], runes[r] = runes[r], runes[l]
					break
				}
			}
		}
		l++
	}
	return string(runes)
}