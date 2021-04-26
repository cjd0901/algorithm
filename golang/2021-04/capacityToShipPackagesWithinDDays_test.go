package _021_04

import (
	"fmt"
	"testing"
)

func TestCapacityToShipPackagesWithinDDays(t *testing.T) {
	days := CapacityToShipPackagesWithinDDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	fmt.Println(days)
}