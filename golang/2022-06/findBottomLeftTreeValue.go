package _022_06

// 找到树左下角的值
// leetcode: https://leetcode.cn/problems/find-bottom-left-tree-value/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindBottomLeftTreeValue(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		tn := q[0]
		q = q[1:]
		if tn.Right != nil {
			q = append(q, tn.Right)
		}
		if tn.Left != nil {
			q = append(q, tn.Left)
		}
		ans = tn.Val
	}
	return ans
}
