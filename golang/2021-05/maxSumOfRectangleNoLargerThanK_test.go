package _021_05

import (
	"fmt"
	"testing"
)

func TestMaxSumOfRectangleNoLargerThanK(t *testing.T) {
	sum := MaxSumOfRectangleNoLargerThanK([][]int{
		{1, 0, 1},
		{0, -2, 3},
	}, 2)
	fmt.Println(sum)
}