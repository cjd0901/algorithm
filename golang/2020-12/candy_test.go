package golang

import (
	"fmt"
	"testing"
)

func TestCandy(t *testing.T) {
	//candy := Candy([]int{1, 0, 2})
	candy := Candy([]int{1, 3, 2, 2, 1})
	//candy := Candy([]int{1, 2, 2})
	fmt.Println(candy)
}