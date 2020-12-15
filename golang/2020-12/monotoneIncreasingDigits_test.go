package golang

import (
	"fmt"
	"testing"
)

func TestMonotoneIncreasingDigits(t *testing.T) {
	digits := MonotoneIncreasingDigits(120)
	fmt.Println(digits)
}