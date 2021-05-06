package _021_05

import (
	"fmt"
	"testing"
)

func TestOnesAndZeros(t *testing.T) {
	subs := OnesAndZeros([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
	fmt.Println(subs)
}