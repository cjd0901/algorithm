package _021_05

import (
	"fmt"
	"testing"
)

func TestSingleElementInASortedArray(t *testing.T) {
	el := SingleElementInASortedArray([]int{3,3,7,7,10,11,11})
	fmt.Println(el)
}