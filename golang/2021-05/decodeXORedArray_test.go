package _021_05

import (
	"fmt"
	"testing"
)

func TestDecodeXORedArray(t *testing.T) {
	array := DecodeXORedArray([]int{6, 2, 7, 3}, 4)
	fmt.Println(array)
}