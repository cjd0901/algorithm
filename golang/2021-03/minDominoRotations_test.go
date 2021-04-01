package _021_03

import (
	"fmt"
	"testing"
)

func TestMinDominoRotations(t *testing.T) {
	sum := MinDominoRotations([]int{1, 2, 3, 4, 6}, []int{6, 6, 6, 6, 5})
	//sum := MinDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4})
	fmt.Println(sum)
}