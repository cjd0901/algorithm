package _021_05

import "math"

// Minimum Number of Days to Make m Bouquets
// leetcode: https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/
// Given an integer array bloomDay, an integer m and an integer k.
// We need to make m bouquets. To make a bouquet, you need to use k adjacent flowers from the garden.
// The garden consists of n flowers, the ith flower will bloom in the bloomDay[i] and then can be used in exactly one bouquet.
// Return the minimum number of days you need to wait to be able to make m bouquets from the garden. If it is impossible to make m bouquets return -1.

func MinimumNumberOfDaysToMakeMBouquets(bloomDay []int, m, k int) int {
	if m * k > len(bloomDay) {
		return -1
	}
	low, high := math.MaxInt32, 0
	n := len(bloomDay)
	for i := 0; i < n; i++ {
		low = min(low, bloomDay[i])
		high = max(high, bloomDay[i])
	}
	var canMake func(days int) bool
	canMake = func(days int) bool {
		bouquets := 0
		flowers := 0
		for i := 0; i < n && bouquets < m; i++ {
			if bloomDay[i] <= days {
				flowers++
				if flowers == k {
					bouquets++
					flowers = 0
				}
			}else {
				flowers = 0
			}
		}
		return bouquets >= m
	}
	for low < high {
		days := (high - low) / 2 + low
		if canMake(days) {
			high = days
		}else {
			low = days + 1
		}
	}
	return low
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}