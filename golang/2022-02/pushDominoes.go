package _022_02

import "bytes"

// Push Dominoes
// leetcode: https://leetcode-cn.com/problems/push-dominoes/
func PushDominoes(dominoes string) string {
	n := len(dominoes)
	q := make([]int, 0)
	time := make([]int, n) // 受力的时间
	for i := range time {
		time[i] = -1
	}
	forces := make([][]byte, n) // 改多米诺牌所受的力
	for i, domino := range dominoes {
		if domino != '.' {
			q = append(q, i)
			time[i]++
			forces[i] = append(forces[i], byte(domino))
		}
	}

	ans := bytes.Repeat([]byte{'.'}, n)
	for len(q) > 0 {
		idx := q[0]
		q = q[1:]
		if len(forces[idx]) > 1 {
			continue
		}
		force := forces[idx][0]
		ans[idx] = force
		ni := idx - 1
		if force == 'R' {
			ni = idx + 1
		}
		if ni >= 0 && ni < n {
			t := time[idx]
			if time[ni] == -1 {
				q = append(q, ni)
				time[ni] = t + 1
				forces[ni] = append(forces[ni], force)
			} else if time[ni] == t + 1 {
				forces[ni] = append(forces[ni], force)
			}
		}
	}
	return string(ans)
}