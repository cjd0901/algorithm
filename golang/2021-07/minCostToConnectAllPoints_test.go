package _021_07

import (
	"fmt"
	"testing"
)

func TestMinCostToConnectAllPoints(t *testing.T) {
	dis := MinCostToConnectAllPoints([][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	})
	fmt.Println(dis)
}
