package dp

import "testing"

func TestCoins(t *testing.T) {
	ans := Coins(5, []float64{0.42, 0.01, 0.42, 0.99, 0.42})
	t.Log(ans)
}