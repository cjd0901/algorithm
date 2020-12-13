package golang

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	//duplicate := ContainsDuplicate([]int{1, 2, 3, 1})
	//duplicate := ContainsDuplicate([]int{1, 2, 3, 1})
	duplicate := ContainsDuplicate([]int{1,1,1,3,3,4,3,2,4,2})
	fmt.Println(duplicate)
}