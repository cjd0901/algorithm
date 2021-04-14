package _021_04

import (
	"fmt"
	"testing"
)

func TestMinimumSizeSubarraySum(t *testing.T) {
	num := MinimumSizeSubarraySum(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Println(num)
}