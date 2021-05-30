package _021_05

import (
	"fmt"
	"testing"
)

func TestNumberOfDiceRollsWithTargetSum(t *testing.T) {
	sum := NumberOfDiceRollsWithTargetSum(2, 5, 6)
	fmt.Println(sum)
}