package _022_06

// 跳跃游戏III
// leetcode: https://leetcode.cn/problems/jump-game-iii/

func JumpGameIII(arr []int, start int) bool {
	canReach := make([]bool, len(arr))
	can := false
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx < 0 || idx > len(arr)-1 || canReach[idx] {
			return
		}
		if arr[idx] == 0 {
			can = true
			return
		}
		canReach[idx] = true
		dfs(idx + arr[idx])
		dfs(idx - arr[idx])
	}
	dfs(start)
	return can
}
