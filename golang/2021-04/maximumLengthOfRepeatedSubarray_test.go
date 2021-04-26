package _021_04

import (
	"fmt"
	"testing"
)

func TestMaximumLengthOfRepeatedSubarray(t *testing.T) {
	length := MaximumLengthOfRepeatedSubarray([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})
	fmt.Println(length)
}