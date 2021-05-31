package _021_05

// Delete Nodes And Return Forest
// leetcode: https://leetcode-cn.com/problems/delete-nodes-and-return-forest/
// Given the root of a binary tree, each node in the tree has a distinct value.
// After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).
// Return the roots of the trees in the remaining forest. You may return the result in any order.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func DeleteNodesAndReturnForest(root *TreeNode, to_delete []int) []*TreeNode {
	deleteMap := make(map[int]bool)
	for _, d := range to_delete {
		deleteMap[d] = true
	}
	tns := []*TreeNode{}
	var dfs func(tn, pre *TreeNode, r bool)
	dfs = func(tn, pre *TreeNode, r bool) {
		if tn == nil {
			return
		}
		dfs(tn.Left, tn, false)
		dfs(tn.Right, tn, true)
		if deleteMap[tn.Val] {
			if tn.Left != nil {
				tns = append(tns, tn.Left)
			}
			if tn.Right != nil {
				tns = append(tns, tn.Right)
			}
			if pre != nil {
				if r {
					pre.Right = nil
				}else {
					pre.Left = nil
				}
			}
		}else {
			if pre == nil {
				tns = append(tns, tn)
			}
		}
	}
	dfs(root, nil, false)
	return tns
}