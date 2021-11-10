package _021_11

// Path Sum II
// leetcode: https://leetcode-cn.com/problems/path-sum-ii/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSumII(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	nodes := make([]int, 0)
	var dfs func(tn *TreeNode, cur, targetSum int)
	dfs = func(tn *TreeNode, cur, targetSum int) {
		if tn == nil {
			return
		}
		cur += tn.Val
		if tn.Left == nil && tn.Right == nil && cur == targetSum {
			nodes = append(nodes, tn.Val)
			ans = append(ans, append([]int(nil), nodes...))
			nodes = nodes[:len(nodes)-1]
			return
		}
		if tn.Left != nil {
			nodes = append(nodes, tn.Val)
			dfs(tn.Left, cur, targetSum)
			nodes = nodes[:len(nodes)-1]
		}
		if tn.Right != nil {
			nodes = append(nodes, tn.Val)
			dfs(tn.Right, cur, targetSum)
			nodes = nodes[:len(nodes)-1]
		}
	}
	dfs(root, 0, targetSum)
	return ans
}

func PathSumII2(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	path := []int{}
	var dfs func(tn *TreeNode, left int)
	dfs = func(tn *TreeNode, left int) {
		if tn == nil {
			return
		}
		left -= tn.Val
		path = append(path, tn.Val)
		defer func() { path = path[:len(path)-1] }()
		if tn.Left == nil && tn.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(tn.Left, left)
		dfs(tn.Right, left)
	}
	dfs(root, targetSum)
	return ans
}