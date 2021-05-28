package _021_05

import (
	"fmt"
	"testing"
)

func TestTotalHammingDistance(t *testing.T) {
	distance := TotalHammingDistance([]int{4, 14, 2})
	fmt.Println(distance)
}