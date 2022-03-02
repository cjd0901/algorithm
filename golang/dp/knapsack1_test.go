package dp

import "testing"

func TestKnapsack1(t *testing.T) {
	ans := Knapsack12(6, 15, [][]int{
		{6, 5},
		{5, 6},
		{6, 4},
		{6, 6},
		{3, 5},
		{7, 2},
	})
	t.Log(ans)
}