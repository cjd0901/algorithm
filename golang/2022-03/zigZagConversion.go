package _022_03

import "strings"

// ZigZag Conversion
// leetcode: https://leetcode-cn.com/problems/zigzag-conversion/
func ZigZagConversion(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	matrix := make([][]byte, numRows)
	diff, row := 0, 0
	for i := 0; i < n; i++ {
		matrix[row] = append(matrix[row], s[i])
		if row == 0 {
			diff = 1
		} else if row == numRows - 1 {
			diff = -1
		}
		row += diff
	}
	sb := strings.Builder{}
	for _, r := range matrix {
		sb.Write(r)
	}
	return sb.String()
}