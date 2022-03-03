package dp

// Longest Path
// https://atcoder.jp/contests/dp/tasks/dp_g
func LongestPath(N int, edges [][]int) int {
	indegrees := make([]int, N+1)
	pedges := make([][]int, N+1)
	for _, edge := range edges {
		pedges[edge[0]] = append(pedges[edge[0]], edge[1])
		indegrees[edge[1]]++
	}
	
	ans := 0
	var dfs func(vertice, path int)
	dfs = func(vertice, path int) {
		ans = max(ans, path)
		for _, pedge := range pedges[vertice] {
			dfs(pedge, path+1)	
		}
	}
	for vertice, indegree := range indegrees {
		if indegree == 0 {
			dfs(vertice, 0)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}