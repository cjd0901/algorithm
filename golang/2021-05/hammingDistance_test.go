package _021_05

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	distance := HammingDistance(1, 4)
	fmt.Println(distance)
}