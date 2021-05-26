package _021_05

import (
	"fmt"
	"testing"
)

func TestReverseSubstringsBetweenEachPairOfParentheses(t *testing.T) {
	s := ReverseSubstringsBetweenEachPairOfParentheses("(ed(et(oc))el)")
	fmt.Println(s)
}