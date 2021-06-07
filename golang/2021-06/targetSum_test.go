package _021_06

import (
	"fmt"
	"testing"
)

func TestTargetSum(t *testing.T) {
	count := TargetSum([]int{1, 1, 1, 1, 1}, 3)
	fmt.Println(count)
}