package _021_08

// Shortest Path Visiting All Nodes
// leetcode: https://leetcode-cn.com/problems/shortest-path-visiting-all-nodes/
func ShortestPathVisitingAllNodes(graph [][]int) int {
	n := len(graph)
	type tuple struct {
		idx, mask, dist int
	}
	q := []tuple{}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, 1<<n)
		visited[i][1<<i] = true
		q = append(q, tuple{i, 1<<i, 0})
	}
	for {
		t := q[0]
		q = q[1:]
		if t.mask == 1<<n-1 {
			return t.dist
		}
		for _, v := range graph[t.idx] {
			m := t.mask | 1<<v
			if !visited[v][m] {
				q = append(q, tuple{v, m, t.dist+1})
				visited[v][m] = true
			}
		}
	}
}