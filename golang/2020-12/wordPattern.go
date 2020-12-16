package golang

import (
	"strings"
)

// Word Pattern
// leetCode: https://leetcode-cn.com/problems/word-pattern/
// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s.

func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	word2ch := make(map[string]byte)
	ch2word := make(map[byte]string)
	if len(pattern) != len(words) {
		return false
	}
	for i, word := range words {
		ch := pattern[i]
		if word2ch[word] > 0 && word2ch[word] != ch || ch2word[ch] != "" && ch2word[ch] != word {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}