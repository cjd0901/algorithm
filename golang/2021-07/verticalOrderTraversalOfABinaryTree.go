package _021_07

import (
	"sort"
)

// Vertical Order Traversal of a Binary Tree
// leetcode: https://leetcode-cn.com/problems/vertical-order-traversal-of-a-binary-tree/
func verticalOrderTraversalOfABinaryTree(root *TreeNode) [][]int {
	m := make(map[int]map[int][]int)
	var dfs func(tn *TreeNode, row, col int)
	dfs = func(tn *TreeNode, row, col int) {
		if tn == nil {
			return
		}
		if _, has := m[col]; !has {
			m[col] = make(map[int][]int)
		}
		m[col][row] = append(m[col][row], tn.Val)
		dfs(tn.Left, row+1, col-1)
		dfs(tn.Right, row+1, col+1)
	}
	dfs(root, 0, 0)
	cols := make([]int, 0)
	for k, _ := range m {
		cols = append(cols, k)
	}
	sort.Ints(cols)
	var ans [][]int
	for _, col := range cols {
		mv := m[col]
		rows := make([]int, 0)
		for k, _ := range mv {
			rows = append(rows, k)
		}
		sort.Ints(rows)
		temp := make([]int, 0)
		for _, row := range rows {
			vals := m[col][row]
			sort.Ints(vals)
			temp = append(temp, vals...)
		}
		ans = append(ans, temp)
	}
	return ans
}