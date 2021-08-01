package _021_08

import (
	"fmt"
	"testing"
)

func TestTheKWeakestRowsInAMatrix(t *testing.T) {
	mat := [][]int{
		{1,1,0,0,0},
		{1,1,1,1,0},
		{1,0,0,0,0},
		{1,1,0,0,0},
		{1,1,1,1,1},
	}
	ans := TheKWeakestRowsInAMatrix(mat, 3)
	fmt.Println(ans)
}