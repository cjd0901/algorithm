package golang

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	//stairs := MinCostClimbingStairs([]int{10, 15, 20})
	stairs := MinCostClimbingStairs([]int{0, 1, 2, 2})
	//stairs := MinCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	fmt.Println(stairs)
}