package _021_04

import (
	"fmt"
	"testing"
)

func TestMinimumPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	path := MinimumPathSum(grid)
	fmt.Println(path)
}