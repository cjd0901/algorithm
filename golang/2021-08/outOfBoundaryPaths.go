package _021_08

// Out of Boundary Paths
// leetcode: https://leetcode-cn.com/problems/out-of-boundary-paths/
func OutOfBoundaryPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	mod := int(1e9+7)
	dis := []struct{x, y int}{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	cache := make([][][]int, m)
	var dfs func(x, y, k int) int
	dfs = func(x, y, k int) int {
		if x == -1 || y == -1 || x == m || y == n {
			return 1
		}
		if k == 0 {
			return 0
		}
		if cache[x][y][k] != -1 {
			return cache[x][y][k]
		}
		ans := 0
		for _, d := range dis {
			ans += dfs(x+d.x, y+d.y, k-1)
			ans %= mod
		}
		cache[x][y][k] = ans
		return ans
	}
	for i := range cache {
		cache[i] = make([][]int, n)
		for j := range cache[i] {
			cache[i][j] = make([]int, maxMove+1)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}
	return dfs(startRow, startColumn, maxMove)
}

func OutOfBoundaryPaths2(m int, n int, maxMove int, startRow int, startColumn int) int {
	mod := int(1e9+7)
	dis := []struct{x, y int}{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	dp := make([][][]int, maxMove+1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	ans := 0
	dp[0][startRow][startColumn] = 1
	for k := 0; k < maxMove; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				count := dp[k][i][j]
				if count > 0 {
					for _, d := range dis {
						i2, j2 := i+d.x, j+d.y
						if i2 >= 0 && j2 >= 0 && i2 < m && j2 < n {
							dp[k+1][i2][j2] = (dp[k+1][i2][j2] + count) % mod
						} else {
							ans = (ans + count) % mod
						}
					}
				}
			}
		}
	}
	return ans
}
