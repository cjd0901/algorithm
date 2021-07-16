package _021_07

import (
	"fmt"
	"testing"
)

func TestFindTheNumberInTheSortedArray(t *testing.T) {
	ans := FindTheNumberInTheSortedArray([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Println(ans)
}