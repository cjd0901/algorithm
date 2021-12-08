package _021_12

// Truncate Sentence
// leetcode: https://leetcode-cn.com/problems/truncate-sentence/

func TruncateSentence(s string, k int) string {
	l, idx, count := len(s), 0, 0
	for i := 0; i <= l; i++ {
		if i == l || s[i] == ' ' {
			count++
			if count == k {
				idx = i
				break
			}
		}
	}
	return s[:idx]
}