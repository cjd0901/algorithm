package _021_05

import (
	"fmt"
	"testing"
)

func TestUncrossedLines(t *testing.T) {
	lines := UncrossedLines([]int{1, 4, 2}, []int{1, 2, 4})
	fmt.Println(lines)
}