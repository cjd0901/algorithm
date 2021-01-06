package _021_01

// Evaluate Division
// leetCode: https://leetcode-cn.com/problems/evaluate-division/

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	// 建图
	type edge struct {
		to int
		weight float64
	}

	graph := make([][]edge, len(id))
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		graph[v] = append(graph[v], edge{w, values[i]})
		graph[w] = append(graph[w], edge{v,1/values[i]})
	}

	bfs := func(start, end int) float64 {
		ratios := make([]float64, len(graph))
		ratios[start] = 1
		queue := []int{start}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if v == end {
				return ratios[v]
			}
			for _, e := range graph[v] {
				if w := e.to; ratios[w] == 0 {
					ratios[w] = ratios[v] * e.weight
					queue = append(queue, w)
				}
			}
		}
		return -1
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if !hasE || !hasS {
			ans[i] = -1
		}else {
			ans[i] = bfs(start, end)
		}
	}
	return ans
}

// wrong answer
func CalcEquationMine(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]float64)
	ans := make([]float64, 0)
	n := len(equations) - 1
	i := 0
	for n >= 0{
		equation := equations[i]
		divisor := equation[0]
		dividend := equation[1]
		if m[divisor] == 0.0 && m[dividend] == 0.0{
			if i == 0 || i >= len(equations) - 1 {
				m[divisor] = 1.0
			}else {
				equations = append(equations, equation)
				values = append(values, values[i])
				n++
			}
		}
		if m[divisor] == 0.0 && m[dividend] != 0.0 {
			m[divisor] = m[dividend] * values[i]
		}
		if m[divisor] != 0.0 && m[dividend] == 0.0 {
			m[dividend] = m[divisor] / values[i]
		}
		i++
		n--
	}
	for _, query := range queries {
		divisor := query[0]
		dividend := query[1]
		if m[divisor] == 0.0 || m[dividend] == 0.0 {
			ans = append(ans, -1.0)
		}else {
			ans = append(ans, m[divisor]/m[dividend])
		}
	}
	return ans
}