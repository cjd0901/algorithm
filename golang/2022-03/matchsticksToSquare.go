package _022_03

import "sort"

// Matchsticks to Square
// leetcode: https://leetcode-cn.com/problems/matchsticks-to-square/
func MatchsticksToSquare(matchsticks []int) bool {
	n := len(matchsticks)
	sum := 0
	for _, matchstick := range matchsticks {
		sum += matchstick
	}
	if sum % 4 != 0 {
		return false
	}
	border := sum / 4
	borders := make([]int, 4)
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == n {
			return borders[0] == borders[1] && borders[1] == borders[2] && borders[2] == borders[3]
		}

		matchstick := matchsticks[index]
		for i := 0; i < 4; i++ {
			if matchstick + borders[i] <= border {
				borders[i] += matchstick
				if dfs(index+1) {
					return true
				}
				borders[i] -= matchstick
			}
		}
		return false
	}
	return dfs(0)
}