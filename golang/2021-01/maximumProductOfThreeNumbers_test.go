package _021_01

import (
	"fmt"
	"testing"
)

func TestMaximumProduct(t *testing.T) {
	product := MaximumProduct([]int{-1, -2, -3, 1, 2, 3})
	fmt.Println(product)
}