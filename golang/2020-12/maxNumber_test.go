package golang

import (
	"fmt"
	"testing"
)

func TestMaxNumber(t *testing.T) {
	array := MaxNumber([]int{6, 7}, []int{9, 3, 9}, 5)
	fmt.Println(array)
}