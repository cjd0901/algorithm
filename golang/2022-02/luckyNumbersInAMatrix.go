package _022_02

// Lucky Numbers in a Matrix
// leetcode: https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/
func LuckyNumbersInAMatrix(matrix [][]int) []int {
	ans := make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	rowMin := make([]int, m)
	colMax := make([]int, n)
	for i, _ := range rowMin {
		rowMin[i] = matrix[i][0]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := matrix[i][j]
			if num < rowMin[i] {
				rowMin[i] = num
			}
			if num > colMax[j] {
				colMax[j] = num
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rowMin[i] == colMax[j] {
				ans = append(ans, rowMin[i])
			}
		}
	}
	return ans
}