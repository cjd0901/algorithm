package golang

import (
	"fmt"
	"testing"
)

func TestMonotoneIncreasingDigits(t *testing.T) {
	//digits := MonotoneIncreasingDigits(1)
	digits := OfficialMonotoneIncreasingDigits(332)
	fmt.Println(digits)
}