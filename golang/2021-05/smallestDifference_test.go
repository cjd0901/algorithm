package _021_05

import (
	"fmt"
	"testing"
)

func TestSmallestDifference(t *testing.T) {
	difference := SmallestDifference([]int{1, 3, 15, 11, 2}, []int{23, 127, 235, 19, 8})
	fmt.Println(difference)
}