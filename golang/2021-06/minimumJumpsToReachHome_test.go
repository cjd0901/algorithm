package _021_06

import (
	"fmt"
	"testing"
)

func TestMinimumJumpsToReachHome(t *testing.T) {
	count := MinimumJumpsToReachHome([]int{1,6,2,14,5,17,4}, 16, 9, 7)
	fmt.Println(count)
}