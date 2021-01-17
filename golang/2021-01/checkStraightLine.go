package _021_01

// Check If It Is a Straight Lien
// leetCode: https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/
// You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point.
// Check if these points make a straight line in the XY plane.

func CheckStraightLine(coordinates [][]int) bool {
	for i := 1; i+1<= len(coordinates) - 1; i++ {
		prev := coordinates[i-1]
		cur := coordinates[i]
		next := coordinates[i+1]
		if (cur[0] - prev[0]) * (next[1] - cur[1]) != (cur[1] - prev[1]) * (next[0] - cur[0]) {
			return false
		}
	}
	return true
}
