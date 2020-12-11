package golang

import (
	"fmt"
	"testing"
)

func TestPredictPartyVictory(t *testing.T) {
	victory := PredictPartyVictory("RRDDD")
	fmt.Println(victory)
}