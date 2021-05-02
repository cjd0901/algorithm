package _021_05

import (
	"fmt"
	"testing"
)

func TestCircusTowerLCCI(t *testing.T) {
	cnt := CircusTowerLCCI([]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 90, 190, 95, 110})
	fmt.Println(cnt)
}