package _021_05

import (
	"fmt"
	"testing"
)

func TestQueensThatCanAttackKing(t *testing.T) {
	queens := QueensThatCanAttackKing([][]int{
		{0, 1},
		{1, 0},
		{4, 0},
		{0, 4},
		{3, 3},
		{2, 4},
	}, []int{0, 0})
	fmt.Println(queens)
}