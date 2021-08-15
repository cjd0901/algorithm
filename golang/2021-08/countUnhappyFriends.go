package _021_08

// Count Unhappy Friends
// leetcode: https://leetcode-cn.com/problems/count-unhappy-friends/
func CountUnhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	idxs := make([][]int, n)
	for i, preference := range preferences {
		idxs[i] = make([]int, n)
		for j, p := range preference {
			idxs[i][p] = j
		}
	}
	pair := make([]int, n)
	for _, p := range pairs {
		pair[p[0]] = p[1]
		pair[p[1]] = p[0]
	}
	ans := 0
	for x, y := range pair {
		i := idxs[x][y]
		for j := 0; j < i; j++ {
			u := preferences[x][j]
			v := pair[u]
			if idxs[u][x] < idxs[u][v] {
				ans++
				break
			}
		}
	}
	return ans
}