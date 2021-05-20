package _021_05

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	ans := TopKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4)
	fmt.Println(ans)
}