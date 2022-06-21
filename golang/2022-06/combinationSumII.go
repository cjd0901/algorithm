package _022_06

import "sort"

// 组合总和II
// leetcode: https://leetcode.cn/problems/combination-sum-ii/

func CombinationSumII(candidates []int, target int) [][]int {
	// 进行排序，后续可以剪枝
	sort.Ints(candidates)

	// a[n][0] 表示数字， a[n][1] 表示该数字出现的次数
	var a [][2]int
	for _, c := range candidates {
		if a == nil || c != a[len(a)-1][0] {
			a = append(a, [2]int{c, 1})
		} else {
			a[len(a)-1][1]++
		}
	}

	ans := make([][]int, 0)
	temp := make([]int, 0)
	var backtrace func(pos, rest int)
	backtrace = func(pos, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}

		if pos == len(a) || rest < a[pos][0] {
			return
		}

		backtrace(pos+1, rest)

		most := min(rest/a[pos][0], a[pos][1])
		for i := 1; i <= most; i++ {
			temp = append(temp, a[pos][0])
			backtrace(pos+1, rest-i*a[pos][0])
		}
		temp = temp[:len(temp)-most]
	}
	backtrace(0, target)
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func CombinationSumII2(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := make([][]int, 0)
	temp := make([]int, 0)
	used := make([]bool, n)
	var backtrace func(idx, rest int)
	backtrace = func(idx, rest int) {
		if rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		if idx == n || used[idx] {
			return
		}
		// 如果不选择当前的数，则跳过进行下一步递归
		backtrace(idx+1, rest)

		// 选择当前的数，需要进行回溯
		used[idx] = true
		temp = append(temp, candidates[idx])
		backtrace(idx+1, rest-candidates[idx])
		temp = temp[:len(temp)-1]
		used[idx] = false
	}
	backtrace(0, target)
	return ans
}
