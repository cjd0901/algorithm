package _021_12

import "unicode"

// Shortest Completing Word
// leetcode:https://leetcode-cn.com/problems/shortest-completing-word/
func ShortestCOmpletingWord(licensePlate string, words []string) string {
	ch := [26]int{}
	for _, c := range licensePlate {
		if unicode.IsLetter(c) {
			ch[unicode.ToLower(c) - 'a']++
		}
	}

	ans := ""
next:
	for _, word := range words {
		c := [26]int{}
		for _, u := range word {
			c[u - 'a']++
		}
		for i := 0; i < 26; i++ {
			if c[i] < ch[i] {
				continue next
			}
		}
		if ans == "" || len(word) < len(ans) {
			ans = word
		}
	}
	return ans
}