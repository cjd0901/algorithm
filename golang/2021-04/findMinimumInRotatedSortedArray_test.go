package _021_04

import (
	"fmt"
	"testing"
)

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	ans := FindMinimumInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2})
	fmt.Println(ans)
}