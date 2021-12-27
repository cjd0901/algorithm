package _021_12

// Even Odd Tree
// leetcode: https://leetcode-cn.com/problems/even-odd-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func EvenOddTree(root *TreeNode) bool {
	i := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		temp := make([]*TreeNode, 0)
		isEven := IsEven(i)
		for len(q) > 0 {
			tn := q[0]
			q = q[1:]
			if len(q) > 0 {
				if (isEven && tn.Val >= q[0].Val) || (!isEven && tn.Val <= q[0].Val) {
					return false
				}
			}
			if (isEven && IsEven(tn.Val)) || (!isEven && !IsEven(tn.Val)) {
				return false
			}
			if tn.Left != nil {
				temp = append(temp, tn.Left)
			}
			if tn.Right != nil {
				temp = append(temp, tn.Right)
			}
		}
		q = temp
		i++
	}
	return true
}

func IsEven(x int) bool {
	return x % 2 == 0
}
