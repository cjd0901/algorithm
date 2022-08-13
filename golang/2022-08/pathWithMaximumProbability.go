package main

type edge struct {
	node     int
	succProb float64
}

func PathWithMaximumProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	visited := make([]bool, n)
	edgeMap := make(map[int][]edge)
	l := len(edges)
	for i := 0; i < l; i++ {
		e := edges[i]
		prob := succProb[i]
		edgeMap[e[0]] = append(edgeMap[e[0]], edge{
			node:     e[1],
			succProb: prob,
		})
		edgeMap[e[1]] = append(edgeMap[e[1]], edge{
			node:     e[0],
			succProb: prob,
		})
	}
	ans := 0.0
	var dfs func(idx int, prob float64)
	dfs = func(idx int, prob float64) {
		if idx == end {
			ans = max(prob, ans)
			return
		}
		visited[idx] = true
		if es, ok := edgeMap[idx]; ok {
			for _, e := range es {
				if !visited[e.node] {
					dfs(e.node, prob*e.succProb)
				}
			}
		}
		visited[idx] = false
	}
	dfs(start, 1.0)
	return ans
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
