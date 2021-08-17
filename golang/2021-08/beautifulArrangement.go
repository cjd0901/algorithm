package _021_08

import "math/bits"

// Beautiful Arrangement
// leetcode: https://leetcode-cn.com/problems/beautiful-arrangement/
func BeautifulArrangement(n int) int {
	visited := make([]bool, n+1)
	match := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}
	ans := 0
	var backtrace func(index int)
	backtrace = func(index int) {
		if index > n {
			ans++
			return
		}
		for _, m := range match[index] {
			if !visited[m] {
				visited[m] = true
				backtrace(index+1)
				visited[m] = false
			}
		}
	}
	backtrace(1)
	return ans
}

func BeautifulArrangement2(n int) int {
	dp := make([]int, 1<<n)
	dp[0] = 1
	for i := 1; i < (1<<n); i++ {
		num := bits.OnesCount(uint(i))
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 && (num%(j+1) == 0 || (j+1)%num == 0) {
				dp[i] += dp[i^(1<<j)]
			}
		}
	}
	return dp[(1<<n)-1]
}