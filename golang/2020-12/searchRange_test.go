package golang

import (
	"fmt"
	"testing"
)

func Test_SearchRange(t *testing.T) {
	//res := SearchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	//res := SearchRange([]int{6}, 6)
	//res := SearchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	res := SearchRange([]int{}, 6)
	fmt.Println(res)
}