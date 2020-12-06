package golang

import (
	"fmt"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	ans := PascalTriangle(5)
	fmt.Println(ans)
}
