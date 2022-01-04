package _022_01

// Flatten a Multilevel Doubly Linked List
// leetcode: https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func FlattenAMutilevelDoublyLinkedList(root *Node) *Node {
	if root == nil {
		return root
	}
	nodes := make([]*Node, 0)
	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		nodes = append(nodes, n)
		dfs(n.Child)
		dfs(n.Next)
	}
	dfs(root)
	nodes[0].Prev = nil
	nodes[len(nodes)-1].Next = nil
	for i := 0; i < len(nodes); i++ {
		if i < len(nodes) - 1 {
			nodes[i].Next = nodes[i+1]
		}
		if i > 0 {
			nodes[i].Prev = nodes[i-1]
		}
		nodes[i].Child = nil
	}
	return nodes[0]
}
