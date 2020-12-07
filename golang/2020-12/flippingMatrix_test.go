package golang

import (
	"fmt"
	"testing"
)

func TestFlippingMatrix(t *testing.T) {
	ans := FlippingMatrix([][]int{
		[]int{0, 0, 1, 1},
		[]int{1, 0, 1, 0},
		[]int{1, 1, 0, 0},
	})
	fmt.Println(ans)
}