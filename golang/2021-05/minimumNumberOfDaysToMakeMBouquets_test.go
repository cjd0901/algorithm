package _021_05

import (
	"fmt"
	"testing"
)

func TestMinimumNumberOfDaysToMakeMBouquets(t *testing.T) {
	days := MinimumNumberOfDaysToMakeMBouquets([]int{1, 10, 3, 10, 2}, 3, 1)
	fmt.Println(days)
}