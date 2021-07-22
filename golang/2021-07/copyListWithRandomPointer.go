package _021_07

// Copy List with Random Pointer
// leetcode: https://leetcode-cn.com/problems/copy-list-with-random-pointer/
type Node struct {
	Val int
	Next *Node
	Random *Node
}

func CopyListWithRandomPointer(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		nodeMap[cur] = &Node{
			Val: cur.Val,
		}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		n := nodeMap[cur]
		n.Next = nodeMap[cur.Next]
		n.Random = nodeMap[cur.Random]
		cur = cur.Next
	}
	return nodeMap[head]
}