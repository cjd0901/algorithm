package _021_05

import (
	"fmt"
	"testing"
)

func TestAllPathFromSourceToTarget(t *testing.T) {
	paths := AllPathFromSourceToTarget([][]int{
		{1, 2},
		{3},
		{3},
		{},
	})
	fmt.Println(paths)
}