package golang

import "fmt"

// Rotate Image
// leetCode: https://leetcode-cn.com/problems/rotate-image/
// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
// DO NOT allocate another 2D matrix and do the rotation.

// 临时矩阵
func RotateImage(matrix [][]int) {
	n := len(matrix)
	temporary := make([][]int, n)
	for i := range temporary {
		temporary[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			temporary[j][n-i-1] = v
		}
	}
	copy(matrix, temporary)
	fmt.Println(matrix)
}

// 水平翻转, 主对角线反转
func RotateImage2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}