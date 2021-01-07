package _021_01

import (
	"fmt"
	"testing"
)

func TestNumberOfProvinces(t *testing.T) {
	provinces := NumberOfProvinces([][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 0, 1, 1},
	})
	fmt.Println(provinces)
}