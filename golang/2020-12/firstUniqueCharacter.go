package golang

// First Unique Character in a String
// leetCode: https://leetcode-cn.com/problems/first-unique-character-in-a-string/
// Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.

func FirstUniqChar(s string) int {
	frequency := make([]int, 26)
	for _, ch := range s {
		frequency[ch - 'a']++
	}
	for index, ch := range s {
		if frequency[ch - 'a'] == 1 {
			return index
		}
	}
	return -1
}
