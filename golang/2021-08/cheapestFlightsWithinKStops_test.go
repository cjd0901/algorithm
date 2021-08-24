package _021_08

import (
	"fmt"
	"testing"
)

func TestCheapestFlightsWithinStops(t *testing.T) {
	ans := CheapestFlightsWithinStops(3, [][]int{
		{0, 1, 2},
		{1, 2, 1},
		{2, 0, 1},
	}, 1, 2, 1)
	fmt.Println(ans)
}