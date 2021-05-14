package _021_05

import (
	"fmt"
	"testing"
)

func TestMostProfitAssigningWork(t *testing.T) {
	profits := MostProfitAssigningWork([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7})
	fmt.Println(profits)
}