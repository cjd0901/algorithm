package _021_12

import (
	"fmt"
	"testing"
)

func TestNumArray(t *testing.T) {
	na := Constructor([]int{1,3,5})
	fmt.Println(na.SumRange(0, 2))
}