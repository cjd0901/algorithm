package _022_06

import "testing"

func TestWiggleSortII(t *testing.T) {
	nums := []int{1, 5, 1, 1, 6, 4}
	WiggleSortII(nums)
	t.Log(nums)
}
