package _021_05

// Queens That Can Attack King
// leetcode: https://leetcode-cn.com/problems/queens-that-can-attack-the-king/
// On an 8x8 chessboard, there can be multiple Black Queens and one White King.
// Given an array of integer coordinates queens that represents the positions of the Black Queens, and a pair of coordinates king that represent the position of the White King, return the coordinates of all the queens (in any order) that can attack the King.

func QueensThatCanAttackKing(queens [][]int, king []int) [][]int {
	ans := [][]int{}
	chessboard := make([][]int, 8)
	for i := 0; i < 8; i++ {
		chessboard[i] = make([]int, 8)
	}
	for _, queen := range queens {
		chessboard[queen[0]][queen[1]] = 1
	}
	var walk func(xStep, yStep int)
	walk = func(xS, yS int) {
		x, y := king[0], king[1]
		for (x + xS <= 7 && x + xS>= 0) && (y + yS <= 7 && y + yS >= 0) {
			x += xS
			y += yS
			if chessboard[x][y] == 1 {
				ans = append(ans, []int{x, y})
				return
			}
		}
	}
	walk(0, 1)
	walk(0, -1)
	walk(1, 1)
	walk(1, -1)
	walk(1, 0)
	walk(-1, 0)
	walk(-1, -1)
	walk(-1, 1)
	return ans
}