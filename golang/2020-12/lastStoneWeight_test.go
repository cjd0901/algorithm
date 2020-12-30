package golang

import (
	"fmt"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	weight := LastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	//weight := LastStoneWeight([]int{2, 2})
	fmt.Println(weight)
}