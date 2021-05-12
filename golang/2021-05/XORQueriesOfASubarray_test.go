package _021_05

import (
	"fmt"
	"testing"
)

func TestXORQueriesOfASubarray(t *testing.T) {
	qs := XORQueriesOfASubarray([]int{1, 3, 4, 8}, [][]int{
		{0, 1},
		{1, 2},
		{0, 3},
		{3, 3},
	})
	fmt.Println(qs)
}