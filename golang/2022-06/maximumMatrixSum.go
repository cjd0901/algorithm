package _022_06

import "math"

// 最大方阵和
// leetcode: https://leetcode.cn/problems/maximum-matrix-sum/
func MaximumMatrixSum(matrix [][]int) int64 {
	negativeCount := 0
	minimum := math.MaxInt32
	sum := int64(0)
	for _, row := range matrix {
		for _, num := range row {
			if num < 0 {
				num = -num
				negativeCount++
			}
			sum += int64(num)
			if num < minimum {
				minimum = num
			}
		}
	}
	if negativeCount%2 == 0 {
		return sum
	}
	return sum - 2*int64(minimum)
}
