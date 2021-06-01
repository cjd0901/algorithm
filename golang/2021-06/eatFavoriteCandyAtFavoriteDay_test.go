package _021_06

import (
	"fmt"
	"testing"
)

func TestEatFavoriteCandyAtFavoriteDay(t *testing.T) {
	ans := EatFavoriteCandyAtFavoriteDay([]int{7, 4, 5, 3, 8}, [][]int{
		{0, 2, 2},
		{4, 2, 4},
		{2, 13, 1000000000},
	})
	fmt.Println(ans)
}