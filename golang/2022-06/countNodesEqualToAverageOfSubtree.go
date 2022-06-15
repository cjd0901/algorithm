package _022_06

// 统计值等于子树平均值的节点数
// leetcode: https://leetcode.cn/problems/count-nodes-equal-to-average-of-subtree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfSubtree(root *TreeNode) int {
	ans := 0
	var dfs func(tn *TreeNode) (int, int)
	dfs = func(tn *TreeNode) (int, int) {
		if tn == nil {
			return 0, 0
		}
		lSum, lCount := dfs(tn.Left)
		rSum, rCount := dfs(tn.Right)
		sum := lSum + rSum + tn.Val
		count := lCount + rCount + 1
		if sum/count == tn.Val {
			ans++
		}
		return sum, count
	}
	dfs(root)
	return ans
}
