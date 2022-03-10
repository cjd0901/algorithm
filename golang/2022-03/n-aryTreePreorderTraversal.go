package _022_03

// N-ary Tree Preorder Traversal
// leetcode: https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
type Node struct {
	Val      int
	Children []*Node
}

func NaryTreePreorderTraversal(root *Node) []int {
	ans := make([]int, 0)
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		ans = append(ans, n.Val)
		if len(n.Children) > 0 {
			for _, child := range n.Children {
				dfs(child)
			}
		}
	}
	dfs(root)
	return ans
}
