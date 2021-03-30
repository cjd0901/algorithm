package _021_03

import (
	"fmt"
	"testing"
)

func TestSearchA2DMatrix(t *testing.T) {
	matrix := [][]int {
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	//matrix := [][]int {
	//	{1},
	//	{3},
	//}
	b := SearchA2DMatrix(matrix, 5)
	fmt.Println(b)
}