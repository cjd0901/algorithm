package _021_05

import (
	"fmt"
	"testing"
)

func TestSatisfiabilityOfEqualityEquations(t *testing.T) {
	ans := SatisfiabilityOfEqualityEquations([]string{"c==c","b==d","x!=z"})
	fmt.Println(ans)
}