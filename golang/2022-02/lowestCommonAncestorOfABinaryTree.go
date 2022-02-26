package _022_02

// Lowest Common Ancestor of a Binary Tree
// leetcode: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestorOfABinaryTree(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestorOfABinaryTree(root.Left, p, q)
	right := LowestCommonAncestorOfABinaryTree(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
