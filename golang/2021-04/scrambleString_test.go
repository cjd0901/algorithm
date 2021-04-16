package _021_04

import (
	"fmt"
	"testing"
)

func TestScrambleString(t *testing.T) {
	ans := ScrambleString("great", "rgeat")
	fmt.Println(ans)
}