package _021_01

import (
	"fmt"
	"testing"
)

func TestCheckStraightLine(t *testing.T) {
	ans := CheckStraightLine([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	})
	//ans := CheckStraightLine([][]int{
	//	{1, 1},
	//	{2, 2},
	//	{3, 4},
	//	{4, 5},
	//	{5, 6},
	//	{7, 7},
	//})
	fmt.Println(ans)
}