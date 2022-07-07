package _022_07

import (
	"strings"
)

// 单词替换
// leetcode: https://leetcode.cn/problems/replace-words/
func ReplaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	dMap := make(map[string]struct{})
	for _, d := range dictionary {
		dMap[d] = struct{}{}
	}
	for i, word := range words {
		for j := 1; j < len(word); j++ {
			if _, ok := dMap[word[0:j]]; ok {
				words[i] = word[0:j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
