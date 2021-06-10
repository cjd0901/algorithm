package _021_06

import (
	"fmt"
	"testing"
)

func TestCoinChange2(t *testing.T) {
	ans := CoinChange2(5, []int{1, 2, 5})
	fmt.Println(ans)
}