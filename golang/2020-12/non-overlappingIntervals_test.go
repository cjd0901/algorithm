package golang

import (
	"fmt"
	"testing"
)

func TestEraseOverLapIntervals(t *testing.T) {
	intervals := EraseOverLapIntervals([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	})
	fmt.Println(intervals)
}