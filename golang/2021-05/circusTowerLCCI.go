package _021_05

import "sort"

// Circus Tower LCCI
// leetcode: https://leetcode-cn.com/problems/circus-tower-lcci/
// A circus is designing a tower routine consisting of people standing atop one anoth­er's shoulders.
// For practical and aesthetic reasons, each person must be both shorter and lighter than the person below him or her.
// Given the heights and weights of each person in the circus, write a method to compute the largest possible number of people in such a tower.

// TODO: 我是笨比
func CircusTowerLCCI(height []int, weight []int) int {
	h2w := make(map[int]int)
	for i := 0; i < len(height); i++ {
		h2w[height[i]] = weight[i]
	}
	sort.Ints(height)
	cnt := 0
	for i, h := range height {
		if i == 0 || h2w[h] > h2w[height[i-1]] {
			cnt++
		}
	}
	return cnt
}