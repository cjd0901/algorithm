package _021_07

import (
	"fmt"
	"testing"
)

func TestGreatestCommonDivisorOfStrings(t *testing.T) {
	ans := GreatestCommonDivisorOfStrings("ABCABCABC", "ABCABC")
	fmt.Println(ans)
}