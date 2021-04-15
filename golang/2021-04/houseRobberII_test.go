package _021_04

import (
	"fmt"
	"testing"
)

func TestHouseRobberII(t *testing.T) {
	//max := HouseRobberII([]int{1, 2, 3, 1})
	max := HouseRobberII([]int{5, 1, 7, 11, 0, 5, 7})
	fmt.Println(max)
}