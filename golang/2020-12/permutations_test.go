package golang

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	permute := Permute([]int{1, 2, 3})
	fmt.Println(permute)
}