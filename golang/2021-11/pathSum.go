package _021_11

// Path Sum
// leetcode: https://leetcode-cn.com/problems/path-sum/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) bool {
	ans := false
	var dfs func(tn *TreeNode, cur, targetSum int)
	dfs = func(tn *TreeNode, cur, targetSum int) {
		if tn == nil {
			return
		}
		cur += tn.Val
		if tn.Left == nil && tn.Right == nil {
			if targetSum == cur {
				ans = true
			}
			return
		}
		if tn.Left != nil {
			dfs(tn.Left, cur, targetSum)
		}
		if tn.Right != nil {
			dfs(tn.Right, cur, targetSum)
		}
	}
	dfs(root, 0, targetSum)
	return ans
}
