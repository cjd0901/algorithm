package golang

import "testing"

func TestRotateImage(t *testing.T) {
	RotateImage2([][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	})
}