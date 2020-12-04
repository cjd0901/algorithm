package golang

import (
	"fmt"
	"testing"
)

func TestSplitArray(t *testing.T) {
	//ans := SplitArray([]int{1, 2, 3, 3, 4, 5})
	ans := SplitArray([]int{1, 2, 3, 3, 4, 4, 5, 5})
	//ans := SplitArray([]int{1, 2, 4, 4, 5})
	//ans := SplitArray([]int{1, 2, 3, 3, 4, 5})
	fmt.Println(ans)
}