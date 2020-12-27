package golang

// Isomorphic Strings
// leetCode: https://leetcode-cn.com/problems/isomorphic-strings/
// Given two strings s and t, determine if they are isomorphic.
// Two strings are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters.
// No two characters may map to the same character but a character may map to itself.

func IsIsomorphic(s, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		m, n := s[i], t[i]
		if s2t[m] > 0 && s2t[m] != n || t2s[n] > 0 && t2s[n] != m {
			return false
		}
		s2t[m] = n
		t2s[n] = m
	}
	return true
}