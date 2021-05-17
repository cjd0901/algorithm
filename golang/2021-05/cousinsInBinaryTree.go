package _021_05

// Cousins in Binary Tree
// leetcode: https://leetcode-cn.com/problems/cousins-in-binary-tree/
// In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
// Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
// We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
// Return true if and only if the nodes corresponding to the values x and y are cousins.

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

func CousinsInBinaryTree(root *TreeNode, x, y int) bool {
	var xPNode, yPNode *TreeNode
	var xDepth, yDepth int
	var xFound, yFound bool
	var dfs func(tn, p *TreeNode, depth int)
	dfs = func(tn, p *TreeNode, depth int) {
		if tn == nil {
			return
		}
		if tn.Val == x {
			xDepth = depth
			xFound = true
			xPNode = p
		}else if tn.Val == y {
			yDepth = depth
			yFound = true
			yPNode = p
		}
		if xFound && yFound {
			return
		}
		dfs(tn.Left, tn, depth+1)
		if xFound && yFound {
			return
		}
		dfs(tn.Right, tn, depth+1)
	}
	dfs(root, nil, 0)
	return xDepth == yDepth && xPNode != yPNode
}