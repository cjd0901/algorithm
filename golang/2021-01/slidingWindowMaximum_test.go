package _021_01

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	//maxSlidingWindow := MaxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	//maxSlidingWindow := MaxSlidingWindow([]int{1}, 1)
	//maxSlidingWindow := MaxSlidingWindow2([]int{1,-1}, 1)
	//maxSlidingWindow := MaxSlidingWindow2([]int{9,11}, 2)
	maxSlidingWindow := MaxSlidingWindow2([]int{-6,-10,-7,-1,-9,9,-8,-4,10,-5,2,9,0,-7,7,4,-2,-10,8,7}, 7)
	fmt.Println(maxSlidingWindow)
}