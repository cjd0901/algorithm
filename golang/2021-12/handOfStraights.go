package _021_12

import (
	"sort"
)

// Hand of Straights
// leetcode: https://leetcode-cn.com/problems/hand-of-straights/
func HandOfStraights(hand []int, W int) bool {
	l := len(hand)
	if l % W != 0 {
		return false
	}

	m := make(map[int]int)
	for _, h := range hand {
		m[h]++
	}

	sort.Ints(hand)
	for _, h := range hand {
		if m[h] == 0 {
			continue
		}
		for i := 0; i < W; i++ {
			if m[h+i] == 0 {
				return false
			}
			m[h+i]--
		}
	}

	return true
}