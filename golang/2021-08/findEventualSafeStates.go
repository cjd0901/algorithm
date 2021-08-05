package _021_08

// Find Eventual Safe States
// leetcode: https://leetcode-cn.com/problems/find-eventual-safe-states/
// 0 => 未被访问的   1 => 已经被访问或存在环中   2 => 可到达终点
func FindEventualSafeStates(graph [][]int) []int {
	n := len(graph)
	color := make([]int, n)
	var check func(i int) bool
	check = func(i int) bool {
		if color[i] != 0 {
			return color[i] == 2
		}
		color[i] = 1
		for _, n := range graph[i] {
			if !check(n) {
				return false
			}
		}
		color[i] = 2
		return true
	}
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if check(i) {
			ans = append(ans, i)
		}
	}
	return ans
}