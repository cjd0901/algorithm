package _021_04

// Capacity To Ship Packages Within D Days
// leetCode: https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
// A conveyor belt has packages that must be shipped from one port to another within D days.

func CapacityToShipPackagesWithinDDays(weights []int, D int) int {
	left, right := 0, 0
	for _, w := range weights {
		if w > left {
			left = w
		}
		right += w
	}
	for left < right {
		mid := (left + right) / 2
		need, sum := 1, 0
		for _, w := range weights {
			if sum + w > mid {
				need ++
				sum = 0
			}
			sum += w
		}
		if need <= D {
			right = mid
		}else {
			left = mid + 1
		}
	}
	return left
}
