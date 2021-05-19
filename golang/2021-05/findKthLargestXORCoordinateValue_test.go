package _021_05

import (
	"fmt"
	"testing"
)

func TestFindKthLargestXORCoordinateValue(t *testing.T) {
	ans := FindKthLargestXORCoordinateValue([][]int{
		{5, 2},
		{1, 6},
	}, 1)
	fmt.Println(ans)
}