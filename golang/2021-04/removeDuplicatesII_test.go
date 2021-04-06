package _021_04

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	//nums := []int{0,0,1,1,1,1,2,3,3}
	nums := []int{0,0,0,0,0}
	l := RemoveDuplicates(nums)
	fmt.Println(l)
	for i := 0; i < l; i++ {
		fmt.Print(nums[i], " ")
	}
}