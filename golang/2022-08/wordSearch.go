package main

import "fmt"

var paths = []struct {
	x, y int
}{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func WordSearch(board [][]byte, word string) bool {
	l := len(word)
	wordBytes := []byte(word)
	var heads [][]int
	for i, col := range board {
		for j, v := range col {
			if wordBytes[0] == v {
				heads = append(heads, []int{i, j})
			}
		}
	}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	var dfs func(i, j, k, length int) bool
	dfs = func(i, j, k, length int) bool {
		if length == l && wordBytes[k] == board[i][j] {
			fmt.Println(11, i, j, k, length)
			return true
		}
		if wordBytes[k] != board[i][j] {
			return false
		}
		visited[i][j] = true
		for _, p := range paths {
			nx := i + p.x
			ny := j + p.y
			if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[0]) && !visited[nx][ny] {
				res := dfs(nx, ny, k+1, length+1)
				if res {
					return true
				}
			}
		}
		visited[i][j] = false
		return false
	}
	for _, head := range heads {
		res := dfs(head[0], head[1], 0, 1)
		if res {
			return true
		}
	}
	return false
}
