package _021_07

import (
	"fmt"
	"testing"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	sum := LargestSumAfterKNegations([]int{
		-8, 3, -5, -3, -5, -2,
	}, 6)
	fmt.Println(sum)
}