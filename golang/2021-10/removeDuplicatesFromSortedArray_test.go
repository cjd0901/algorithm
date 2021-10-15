package _021_10

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	nums := []int{1,1,2}
	ans := RemoveDuplicatesFromSortedArray(nums)
	fmt.Println(ans)
	fmt.Println(nums)
}