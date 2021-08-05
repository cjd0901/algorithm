package _021_08

import (
	"fmt"
	"testing"
)

func TestFindEventualSafeStates(t *testing.T) {
	ans := FindEventualSafeStates([][]int{
		{1,2},
		{2,3},
		{5},
		{0},
		{5},
		{},
		{},
	})
	fmt.Println(ans)
}