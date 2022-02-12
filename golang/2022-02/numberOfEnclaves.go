package _022_02

// Number of Enclaves
// leetcode: https://leetcode-cn.com/problems/number-of-enclaves/

func NumberOfEnclaves(grid [][]int) int {
	paths := []struct{ x, y int }{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0
		for _, path := range paths {
			nx, ny := x + path.x, y + path.y
			if nx < m && nx >= 0 && ny < n && ny >= 0 {
				dfs(nx, ny)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}