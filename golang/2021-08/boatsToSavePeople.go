package _021_08

import (
	"sort"
)

// Boats to Save People
// leetcode: https://leetcode-cn.com/problems/boats-to-save-people/
func BoatsToSavePeople(people []int, limit int) int {
	sort.Ints(people)
	ans := 0
	l, r := 0, len(people)-1
	for l <= r {
		if people[r] + people[l] <= limit {
			ans++
			l++
			r--
		} else {
			ans++
			r--
		}
	}
	return ans
}