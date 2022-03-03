package dp

// Grid 1
// https://atcoder.jp/contests/dp/tasks/dp_h
func Grid1(H, W int, grid [][]byte) int {
	dp := make([][]int, H)
	for i := range dp {
		dp[i] = make([]int, W)
	}
	mod := int(1e9 + 7)

	for i := 0; i < W; i++ {
		if grid[0][i] == '#' {
			break
		}
		dp[0][i] = 1
	}

	for i := 0; i < H; i++ {
		if grid[i][0] == '#' {
			break
		}
		dp[i][0] = 1
	}

	for i := 1; i < H; i++ {
		for j := 1; j < W; j++ {
			value := grid[i][j]
			if value == '.' {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod
			}
		}
	}
	return dp[H-1][W-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
