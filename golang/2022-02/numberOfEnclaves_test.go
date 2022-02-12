package _022_02

import (
	"fmt"
	"testing"
)

func TestNumberOfEnclaves(t *testing.T) {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	ans := NumberOfEnclaves(grid)
	fmt.Println(ans)
}