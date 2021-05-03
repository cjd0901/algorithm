package _021_05

import (
	"fmt"
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	winner := PredictTheWinner([]int{1, 5, 2})
	fmt.Println(winner)
}