package _021_12

// House Robber III
// leetcode: https://leetcode-cn.com/problems/house-robber-iii/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HouseRobberIII(root *TreeNode) int {
	f, g := make(map[*TreeNode]int), make(map[*TreeNode]int)
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		dfs(tn.Right)
		f[tn] = tn.Val + g[tn.Left] + g[tn.Right]
		g[tn] = max(f[tn.Left], g[tn.Left]) + max(f[tn.Right], g[tn.Right])
	}
	dfs(root)
	return max(f[root], g[root])
}

func HouseRobberIII2(root *TreeNode) int {
	var dfs func(tn *TreeNode) (int, int)
	dfs = func(tn *TreeNode) (int, int) {
		if tn == nil {
			return 0, 0
		}
		lf, lg := dfs(tn.Left)
		rf, rg := dfs(tn.Right)
		l := tn.Val + lg + rg
		g := max(lf, lg) + max(rf, rg)
		return l, g
	}
	l, g := dfs(root)
	return max(l, g)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}