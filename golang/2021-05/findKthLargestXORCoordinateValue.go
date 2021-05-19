package _021_05

import "sort"

// Find Kth Largest XOR Coordinate Value
// leetcode: https://leetcode-cn.com/problems/find-kth-largest-xor-coordinate-value/

func FindKthLargestXORCoordinateValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	result := []int{}
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, n := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ n
			result = append(result, pre[i+1][j+1])
		}
	}
	sort.Ints(result)
	return result[m*n-k]
}