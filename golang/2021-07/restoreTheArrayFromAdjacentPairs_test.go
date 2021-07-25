package _021_07

import (
	"fmt"
	"testing"
)

func TestRestoreTheArrayFromAdjacentPairs(t *testing.T) {
	ans := RestoreTheArrayFromAdjacentPairs([][]int{
		{2, 1},
		{3, 4},
		{3, 2},
	})
	fmt.Println(ans)
}