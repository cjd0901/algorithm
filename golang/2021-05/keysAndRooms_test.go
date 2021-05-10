package _021_05

import (
	"fmt"
	"testing"
)

func TestKeysAndRooms(t *testing.T) {
	open := KeysAndRooms([][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	})
	fmt.Println(open)
}