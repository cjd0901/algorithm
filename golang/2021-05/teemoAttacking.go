package _021_05

// TeemoAttacking
// leetcode: https://leetcode-cn.com/problems/teemo-attacking/
// Return the total time that the enemy is in a poisoned condition.

func TeemoAttacking(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 {
		return 0
	}
	total := 0
	for i := 0; i < n - 1; i++ {
		total += min(timeSeries[i + 1] - timeSeries[i], duration)
	}
	return total + duration
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
