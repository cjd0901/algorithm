package _021_01

import (
	"fmt"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	//by5 := PrefixesDivBy5([]int{0, 1, 1})
	//by5 := PrefixesDivBy5([]int{1, 1, 1})
	by5 := PrefixesDivBy5([]int{0, 1, 1, 1, 1, 1})
	fmt.Println(by5)
}