package _021_03

// Search a 2D Matrix
// leetCode: https://leetcode-cn.com/problems/search-a-2d-matrix/
// Write an efficient algorithm that searches for a value in an m x n matrix.
// This matrix has the following properties:
// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.

func SearchA2DMatrix(matrix [][]int, target int) bool {
	left := 0
	colNum := len(matrix)
	rowNum := len(matrix[0])
	right :=  colNum * rowNum - 1
	for left <= right {
		mid := (left + right) / 2
		col := mid / rowNum
		row := mid % rowNum
		midValue := matrix[col][row]
		if target == midValue {
			return true
		}else if target < midValue {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}
	return false
}
