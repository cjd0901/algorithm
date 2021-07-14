package _021_07

import (
	"fmt"
	"testing"
)

func TestMinimumAbsoluteSumDifference(t *testing.T) {
	ans := MinimumAbsoluteSumDifference([]int{1,10,4,4,2,7}, []int{9,3,5,1,7,4})
	fmt.Println(ans)
}