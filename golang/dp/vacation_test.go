package dp

import "testing"

func TestVacation(t *testing.T) {
	ans := Vacation(7, [][]int{
		{6, 7, 8},
		{8, 8, 3},
		{2, 5, 2},
		{7, 8, 6},
		{4, 6, 8},
		{2, 3, 4},
		{7, 5, 1},
	})
	t.Log(ans)
}