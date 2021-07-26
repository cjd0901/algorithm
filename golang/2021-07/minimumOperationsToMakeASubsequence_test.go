package _021_07

import (
	"fmt"
	"testing"
)

func TestMinimumOperationsToMakeASubsequence(t *testing.T) {
	ans := MinimumOperationsToMakeASubsequence([]int{
		6,4,8,1,3,2,
	}, []int{
		4,7,6,2,3,8,6,1,
	})
	fmt.Println(ans)
}