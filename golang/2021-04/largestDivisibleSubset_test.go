package _021_04

import (
	"fmt"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	subset := LargestDivisibleSubset([]int{3, 4, 16, 8})
	fmt.Println(subset)
}