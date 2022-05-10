package _022_04

import "testing"

func TestFindKthLargest(t *testing.T) {
	ans := FindKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	t.Log(ans)
}
