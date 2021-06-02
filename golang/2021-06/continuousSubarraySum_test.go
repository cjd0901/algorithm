package _021_06

import (
	"fmt"
	"testing"
)

func TestContinuousSubarraySum(t *testing.T) {
	sum := ContinuousSubarraySum([]int{23,2,6,4,7}, 6)
	fmt.Println(sum)
}