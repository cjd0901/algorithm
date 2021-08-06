package _021_08

import (
	"fmt"
	"testing"
)

func TestShortestPathVisitingAllNodes(t *testing.T) {
	ans := ShortestPathVisitingAllNodes([][]int{
		{1,2,3},
		{0},
		{0},
		{0},
	})
	fmt.Println(ans)
}