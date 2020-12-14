package golang

import (
	"sort"
)

// Group Anagrams
// leetCode: https://leetcode-cn.com/problems/group-anagrams/
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

func sortStr(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func GroupAnagrams(strs []string) (ans [][]string) {
	sMap := make(map[string][]string)
	for _, s := range strs {
		sortS := sortStr(s)
		sMap[sortS] = append(sMap[sortS], s)
	}
	for _, list := range sMap {
		ans = append(ans, list)
	}
	return
}

