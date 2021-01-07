package _021_01

// Number of Provinces
// leetCode: https://leetcode-cn.com/problems/number-of-provinces/
// There are n cities. Some of them are connected, while some are not.
// If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.
// A province is a group of directly or indirectly connected cities and no other cities outside of the group.
// You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.
// Return the total number of provinces.

// 并查集
func NumberOfProvinces(isConnected [][]int) int {
	n := len(isConnected)
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	var find func(city int) int
	find = func(city int) int {
		if city == fa[city] {
			return city
		}else {
			fa[city] = find(fa[city])
			return fa[city]
		}
	}
	union := func(i, j int) {
		fa[find(j)] = find(i)
	}
	count := 0
	for i, c1 := range isConnected{
		for j, c2 := range c1 {
			if c2 == 1 {
				union(i, j)
			}
		}
	}
	for i, p := range fa {
		if i == p {
			count ++
		}
	}
	return count
}

// DFS
func NumberOfProvinces2(isConnected [][]int) int {
	visit := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		visit[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !visit[to] {
				dfs(to)
			}
		}
	}
	count := 0
	for i, v := range visit {
		if !v {
			count++
			dfs(i)
		}
	}
	return count
}

// BFS
func NumberOfProvinces3(isConnected [][]int) int {
	visit := make([]bool, len(isConnected))
	count := 0
	for i, v := range visit {
		if !v {
			count++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				visit[from] = true
				for to, conn := range isConnected[from] {
					if conn == 1 && !visit[to] {
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return count
}