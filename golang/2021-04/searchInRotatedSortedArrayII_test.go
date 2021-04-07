package _021_04

import (
	"fmt"
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	nums := []int{2,5,6,0,0,1,2}
	ans := SearchInRotatedSortedArray(nums, 5)
	fmt.Println(ans)
}