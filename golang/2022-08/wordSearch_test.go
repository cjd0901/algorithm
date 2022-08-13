package main

import "testing"

func TestWordSearch(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCESCFSADEE"
	t.Log(WordSearch(board, word))
}
