package _021_08

import (
	"fmt"
	"testing"
)

func TestCountUnhappyFriends(t *testing.T) {
	ans := CountUnhappyFriends(4, [][]int{
		{1, 2, 3},
		{3, 2, 0},
		{3, 1, 0},
		{1, 2, 0},
	}, [][]int{
		{0, 1},
		{2, 3},
	})
	fmt.Println(ans)
}