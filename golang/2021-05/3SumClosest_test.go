package _021_05

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	sum := ThreeSumClosest([]int{-1, 2, 1, -4}, 1)
	fmt.Println(sum)
}