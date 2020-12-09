package golang

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	paths := UniquePaths(3, 7)
	fmt.Println(paths)
}