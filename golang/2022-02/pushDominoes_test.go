package _022_02

import "testing"

func TestPushDominoes(t *testing.T) {
	dominoes := ".L.R...LR..L.."
	ans := PushDominoes(dominoes)
	t.Log(ans)
}