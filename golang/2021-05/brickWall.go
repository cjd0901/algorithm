package _021_05

// Brick Wall
// leetcode: https://leetcode-cn.com/problems/brick-wall/
// There is a rectangular brick wall in front of you with n rows of bricks. The ith row has some number of bricks each of the same height (i.e., one unit) but they can be of different widths. The total width of each row is the same.

func BrickWall(wall [][]int) int {
	wallM := make(map[int]int)
	for _, bricks := range wall {
		sum := 0
		for _, brick := range bricks[:len(bricks) - 1] {
			sum += brick
			wallM[sum]++
		}
	}
	maxBorder := 0
	for _, v := range wallM {
		if v > maxBorder {
			maxBorder = v
		}
	}
	return len(wall) - maxBorder
}