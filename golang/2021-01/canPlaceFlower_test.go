package _021_01

import (
	"fmt"
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	flowers := CanPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)
	fmt.Println(flowers)
}