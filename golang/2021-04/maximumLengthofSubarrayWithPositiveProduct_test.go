package _021_04

import (
	"fmt"
	"testing"
)

func TestMaximumLengthOfSubarrayWithPositiveProduct(t *testing.T) {
	l := MaximumLengthOfSubarrayWithPositiveProduct([]int{1, -2, -3, 4})
	fmt.Println(l)
}