package _021_01

import (
	"fmt"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	ranges := SummaryRanges([]int{0, 1, 2, 4, 5, 7})
	fmt.Println(ranges)
}