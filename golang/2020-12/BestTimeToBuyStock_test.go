package golang

import (
	"fmt"
	"testing"
)

func TestMaxStockProfit(t *testing.T) {
	//profit := MaxStockProfit([]int{1, 3, 2, 8, 4, 9}, 2)
	profit := MaxStockProfit2([]int{1, 3, 2, 8, 4, 9}, 2)
	fmt.Println(profit)
}