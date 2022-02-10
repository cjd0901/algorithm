package _022_02

// Path with Maximum Gold
// leetcode: https://leetcode-cn.com/problems/path-with-maximum-gold/
func PathWithMaximumGold(grid [][]int) int {
	paths := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := 0

	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		gold += grid[x][y]
		if gold > ans {
			ans = gold
		}

		temp := grid[x][y]
		grid[x][y] = 0
		for _, path := range paths {
			nx, ny := x + path.x, y + path.y
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[x]) && grid[nx][ny] > 0 {
				dfs(nx, ny, gold)
			}
		}

		grid[x][y] = temp
	}

	for i, row := range grid {
		for j, gold := range row {
			if gold > 0 {
				dfs(i, j , 0)
			}
		}
	}
	return ans
}