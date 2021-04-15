package _021_04

import (
	"fmt"
	"testing"
)

func TestHouseRobber(t *testing.T) {
	ans := HouseRobber([]int{5, 1, 7, 11, 0, 5, 7})
	fmt.Println(ans)
}