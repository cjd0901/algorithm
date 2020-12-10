package golang

import (
	"fmt"
	"testing"
)

func TestLemonadeChange(t *testing.T) {
	//change := LemonadeChange([]int{5, 5, 5, 10, 20})
	//change := LemonadeChange([]int{5,5,10})
	//change := LemonadeChange([]int{10, 10})
	//change := LemonadeChange([]int{5, 5, 10, 10, 20})
	change := LemonadeChange([]int{5,5,5,10,5,20,5,10,5,20})

	fmt.Println(change)
}