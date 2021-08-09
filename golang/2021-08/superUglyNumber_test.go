package _021_08

import (
	"fmt"
	"testing"
)

func TestSuperUglyNumber(t *testing.T) {
	ans := SuperUglyNumber(12, []int{2,7,13,19})
	fmt.Println(ans)
}