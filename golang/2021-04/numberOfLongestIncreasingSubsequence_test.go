package _021_04

import (
	"fmt"
	"testing"
)

func TestNumberOfLongestIncreasingSubsequence(t *testing.T) {
	num := NumberOfLongestIncreasingSubsequence([]int{1, 3, 5, 4, 7, 2, 3, 5})
	fmt.Println(num)
}