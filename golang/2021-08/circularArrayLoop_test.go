package _021_08

import (
	"fmt"
	"testing"
)

func TestCircularArrayLoop(t *testing.T) {
	ans := CircularArrayLoop([]int{2,-1,1,2,2})
	fmt.Println(ans)
}