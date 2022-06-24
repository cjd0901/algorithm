package _022_06

import (
	"testing"
)

func TestTwoCityScheduling(t *testing.T) {
	ans := TwoCityScheduling([][]int{
		{259, 770},
		{448, 54},
		{926, 667},
		{184, 139},
		{840, 118},
		{577, 469},
	})
	t.Log(ans)
}
