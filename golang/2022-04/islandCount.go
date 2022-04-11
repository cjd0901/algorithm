package _022_04

// nowcoder: https://www.nowcoder.com/practice/0c9664d1554e466aa107d899418e814e?tpId=117&&tqId=37833&&companyId=239&rp=1&ru=/company/home/code/239&qru=/ta/job-code-high/question-ranking
type path struct {
	x, y int
}

var paths = []path{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func islandCount(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x == m || y == n || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		for _, p := range paths {
			nx := x + p.x
			ny := y + p.y
			dfs(nx, ny)
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
