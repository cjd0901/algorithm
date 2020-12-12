package golang

import (
	"fmt"
	"testing"
)

func TestWiggleSubsequence(t *testing.T) {
	//subsequence := WiggleSubsequence([]int{1,17,5,10,13,15,10,5,16,8})
	//subsequence := WiggleSubsequence([]int{1, 7, 4, 9, 2, 5})
	//subsequence := WiggleSubsequence([]int{1,2,3,4,5,6,7,8,9})
	//subsequence := WiggleSubsequence([]int{})
	//subsequence := WiggleSubsequence([]int{1})
	//subsequence := WiggleSubsequence([]int{0,0})
	subsequence := WiggleSubsequence([]int{3,3,3,2,5})
	fmt.Println(subsequence)
}