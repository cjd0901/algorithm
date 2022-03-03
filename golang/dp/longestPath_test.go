package dp

import (
	"fmt"
	"testing"
)

func TestLongestPath(t *testing.T) {
	ans := LongestPath(5, [][]int{
		{5, 3},
		{2, 3},
		{2, 4},
		{5, 2},
		{5, 1},
		{1, 4},
		{4, 3},
		{1, 3},
	})
	fmt.Println(ans)
}