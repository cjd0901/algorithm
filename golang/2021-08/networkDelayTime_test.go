package _021_08

import (
	"fmt"
	"testing"
)

func TestNetworkDelayTime(t *testing.T) {
	ans := NetworkDelayTime([][]int{
		{2,1,1},
		{2,3,1},
		{3,4,1},
	},4,2)
	fmt.Println(ans)
}