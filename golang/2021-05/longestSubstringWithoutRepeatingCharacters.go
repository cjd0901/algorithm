package _021_05

// Longest Substring Without Repeating Characters
// leetcode: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// Given a string s, find the length of the longest substring without repeating characters.

func LongestSubstringWithoutRepeatingCharacters(s string) int {
	l := 0
	chIndex := make(map[byte]int)
	start := 0
	for end := 0; end < len(s); end++ {
		ch := s[end]
		if index, ok := chIndex[ch]; ok {
			start = max(index+1, start)
		}
		l = max(l, end - start + 1)
		chIndex[ch] = end
	}
	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}