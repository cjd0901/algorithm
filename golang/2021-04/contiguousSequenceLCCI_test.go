package _021_04

import (
	"fmt"
	"testing"
)

func TestContiguousSequenceLCCI(t *testing.T) {
	sum := ContiguousSequenceLCCI([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(sum)
}