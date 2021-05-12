package _021_05

import (
	"fmt"
	"testing"
)

func TestMaxNumberOfKSumPairs(t *testing.T) {
	max := MaxNumberOfKSumPairs([]int{1, 2, 3, 4}, 5)
	fmt.Println(max)
}