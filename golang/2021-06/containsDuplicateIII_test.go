package _021_06

import (
	"fmt"
	"testing"
)

func TestContainsDuplicateIII(t *testing.T) {
	ans := ContainsDuplicateIII([]int{1,5,9,1,5,9}, 2, 3)
	fmt.Println(ans)
}