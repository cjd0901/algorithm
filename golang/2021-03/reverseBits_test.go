package _021_03

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	var input uint32
	input = 43261596
	bits := ReverseBits(input)
	fmt.Println(bits)
}