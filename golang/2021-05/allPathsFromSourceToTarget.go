package _021_05

// All Paths From Source to Target
// leetcode: https://leetcode-cn.com/problems/all-paths-from-source-to-target/
// Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths from node 0 to node n - 1, and return them in any order.
// The graph is given as follows: graph[i] is a list of all nodes you can visit from node i (i.e., there is a directed edge from node i to node graph[i][j]).

func AllPathFromSourceToTarget(graph [][]int) [][]int {
	ans := make([][]int, 0)
	n := len(graph) - 1
	temp := make([]int, 0)
	var dfs func(node int)
	dfs = func(node int) {
		temp = append(temp, node)
		if node == n {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		if len(graph[node]) == 0 {
			return
		}
		for _, no := range graph[node] {
			dfs(no)
			temp = temp[:len(temp) - 1]
		}
	}
	dfs(0)
	return ans
}