package _021_08

import (
	"fmt"
	"testing"
)

func TestMaximumNumberOfEatenApples(t *testing.T) {
	ans := MaximumNumberOfEatenApples([]int{9, 2}, []int{3,5})
	fmt.Println(ans)
}