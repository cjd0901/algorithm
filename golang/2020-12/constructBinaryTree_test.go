package golang

import "testing"

func TestBuildTree(t *testing.T) {
	BuildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
}