package _021_05

import (
	"fmt"
	"testing"
)

func TestMinimumElementsToAddToFormAGivenSum(t *testing.T) {
	elements := MinimumElementsToAddToFormAGivenSum([]int{1, -1, 1}, 3, -4)
	fmt.Println(elements)
}