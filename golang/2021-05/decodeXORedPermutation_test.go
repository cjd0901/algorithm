package _021_05

import (
	"fmt"
	"testing"
)

func TestDecodeXORedPermutation(t *testing.T) {
	perm := DecodeXORedPermutation([]int{6,5,4,6})
	fmt.Println(perm)
}