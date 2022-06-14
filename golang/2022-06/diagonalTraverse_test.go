package _022_06

import "testing"

func TestDiagonalTraverse(t *testing.T) {
	ans := DiagonalTraverse([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	t.Log(ans)
}
