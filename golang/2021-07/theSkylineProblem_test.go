package _021_07

import (
	"fmt"
	"testing"
)

func TestTheSkylineProblem(t *testing.T) {
	ans := TheSkylineProblem([][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	})
	fmt.Println(ans)
}