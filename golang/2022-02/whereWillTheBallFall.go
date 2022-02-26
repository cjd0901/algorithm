package _022_02

// Where Will the Ball Fall
// leetcode: https://leetcode-cn.com/problems/where-will-the-ball-fall/
func WhereWillTheBallFall(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)
	for i := range ans {
		col := i
		for _, row := range grid {
			dir := row[col]
			col += dir
			if col < 0 || col == n || row[col] != dir {
				col = -1
				break
			}
		}
		ans[i] = col
	}
	return ans
}