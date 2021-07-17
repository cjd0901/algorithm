package _021_07

import "sort"

// Min Cost to Connect All Points
// leetcode: https://leetcode-cn.com/problems/min-cost-to-connect-all-points/
type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
		rank[i] = i
	}
	return &unionFind{
		parent,
		 rank,
	}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.rank[x] < uf.rank[y] {
		fx, fy = fy, fx
	}
	uf.rank[fx] += uf.rank[fy]
	uf.parent[fy] = fx
	return true
}

func distance(a, b []int) int {
	return abs(a[0] - b[0]) + abs(a[1] - b[1])
}

func MinCostToConnectAllPoints(points [][]int) int {
	n := len(points)
	ans := 0
	type edge struct {
		v, w, dis int
	}
	var edges []edge
	for i, p := range points {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{
				v:   i,
				w:   j,
				dis: distance(p, points[j]),
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})

	uf := newUnionFind(n)
	left := n-1
	for _, edge := range edges {
		if uf.union(edge.v, edge.w) {
			ans += edge.dis
			left--
			if left == 0 {
				return ans
			}
		}
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}