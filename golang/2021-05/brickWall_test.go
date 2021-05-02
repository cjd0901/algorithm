package _021_05

import (
	"fmt"
	"testing"
)

func TestBrickWall(t *testing.T) {
	wall := BrickWall([][]int{
		{1, 2, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{2, 4},
		{3, 1, 2},
		{1, 3, 1, 1},
	})
	fmt.Println(wall)
}