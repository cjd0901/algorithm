package _021_05

// Successor
// leetcode: https://leetcode-cn.com/problems/successor-lcci/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Successor(root, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	if p.Val >= root.Val {
		return Successor(root.Right, p)
	}else {
		left := Successor(root.Left, p)
		if left == nil {
			return root
		}
		return left
	}
}

func Successor3(root, p *TreeNode) *TreeNode {
	var pre *TreeNode
	for root.Val != p.Val {
		if root.Val < p.Val {
			root = root.Right
		}else {
			pre = root
			root = root.Left
		}
	}
	if root.Right == nil {
		return pre
	} else {
		root = root.Right
		for root.Left != nil {
			root = root.Left
		}
		return root
	}
}

func Successor2(root, p *TreeNode) *TreeNode {
	tnList := make([]*TreeNode, 0)
	var dfs func(tr *TreeNode)
	dfs = func(tr *TreeNode) {
		if tr == nil {
			return
		}
		dfs(tr.Left)
		tnList = append(tnList, tr)
		dfs(tr.Right)
	}
	dfs(root)
	for i, tn := range tnList {
		if tn == p && i + 1 < len(tnList) {
			return tnList[i+1]
		}
	}
	return nil
}