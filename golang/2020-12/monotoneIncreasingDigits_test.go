package golang

import (
	"fmt"
	"testing"
)

func TestMonotoneIncreasingDigits(t *testing.T) {
	digits := MonotoneIncreasingDigits(1)
	fmt.Println(digits)
}