package _022_06

import "testing"

func TestDuplicateZeros(t *testing.T) {
	arr := []int{0, 0, 0, 0, 0, 0, 0}
	DuplicateZeros(arr)
	t.Log(arr)
}
