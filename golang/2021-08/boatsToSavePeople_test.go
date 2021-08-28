package _021_08

import (
	"fmt"
	"testing"
)

func TestBoatsToSavePeople(t *testing.T) {
	ans := BoatsToSavePeople([]int{3, 2, 2, 1}, 3)
	fmt.Println(ans)
}