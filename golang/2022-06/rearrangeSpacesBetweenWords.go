package _022_06

import (
	"strings"
)

// 新排列单词间的空格
// leetcode: https://leetcode.cn/problems/rearrange-spaces-between-words/
func RearrangeSpacesBetweenWords(text string) string {
	cnt := strings.Count(text, " ")
	words := strings.Fields(text)
	l := len(words)
	if l > 1 {
		return strings.Join(words, strings.Repeat(" ", cnt/(l-1))) + strings.Repeat(" ", cnt%(l-1))
	}
	return strings.Join(words, "") + strings.Repeat(" ", cnt)
}
