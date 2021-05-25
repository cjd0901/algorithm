package _021_05

// Satisfiability of Equality Equations
// leetcode: https://leetcode-cn.com/problems/satisfiability-of-equality-equations/

func SatisfiabilityOfEqualityEquations(equations []string) bool {
	fa := make(map[byte]byte)
	_ = make(map[byte]byte)
	var find func(x byte) byte
	var merge func(x, y byte)
	find = func(x byte) byte {
		if fa[x] == x {
			return x
		}else {
			fa[x] = find(fa[x])
			return fa[x]
		}
	}
	merge = func(x, y byte) {
		fa[find(y)] = find(x)
	}
	for _, equation := range equations {
		x := equation[0]
		y := equation[3]
		if _, ok := fa[x]; !ok {
			fa[x] = x
		}
		if _, ok := fa[y]; !ok {
			fa[y] = y
		}
		symbol := equation[1]
		if symbol == '=' {
			merge(x, y)
		}
	}
	for _, equation := range equations {
		x := equation[0]
		y := equation[3]
		symbol := equation[1]
		if symbol != '=' {
			if find(x) == find(y) {
				return false
			}
		}
	}
	return true
}