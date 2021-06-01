package _021_06

// Minimum Jumps To Reach Home
// leetcode: https://leetcode-cn.com/problems/minimum-jumps-to-reach-home/

func MinimumJumpsToReachHome(forbidden []int, a int, b int, x int) int {
	count := -1
	visited := make([]bool, 8000)
	for _, f := range forbidden {
		visited[f] = true
	}
	var dfs func(cur, curCount int, back bool)
	dfs = func(cur, curCount int, back bool) {
		if cur > 6000 {
			return
		}
		if cur == x {
			count = curCount
			return
		}
		if !visited[cur+a] {
			visited[cur+a] = true
			dfs(cur + a, curCount+1, false)
		}
		if !back && cur-b >= 0 && !visited[cur-b] {
			dfs(cur - b, curCount+1, true)
		}
	}
	dfs(0, 0, true)
	return count
}