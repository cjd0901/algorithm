package golang

import (
	"fmt"
	"testing"
)

func TestAssignCookies(t *testing.T) {
	cookies := AssignCookies([]int{1, 2, 3}, []int{1, 1})
	//cookies := AssignCookies([]int{1, 2}, []int{1, 2, 3})
	fmt.Println(cookies)
}