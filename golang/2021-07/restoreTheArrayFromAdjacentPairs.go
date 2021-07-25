package _021_07

// Restore the Array From Adjacent Pairs
// leetcode: https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs/
func RestoreTheArrayFromAdjacentPairs(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	ans := make([]int, n)
	adjs := make(map[int][]int)
	for _, p := range adjacentPairs {
		adjs[p[0]] = append(adjs[p[0]], p[1])
		adjs[p[1]] = append(adjs[p[1]], p[0])
	}
	for k, adj := range adjs {
		if len(adj) == 1 {
			ans[0] = k
			break
		}
	}
	ans[1] = adjs[ans[0]][0]
	for i := 2; i < n; i++ {
		pre := ans[i-1]
		if ans[i-2] == adjs[pre][0] {
			ans[i] = adjs[pre][1]
		} else {
			ans[i] = adjs[pre][0]
		}
	}
	return ans
}