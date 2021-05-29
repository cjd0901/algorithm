package _021_05

import (
	"fmt"
	"testing"
)

func TestNumberOfSubmatricesThatSumToTarget(t *testing.T) {
	count := NumberOfSubmatricesThatSumToTarget([][]int{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}, 0)
	fmt.Println(count)
}