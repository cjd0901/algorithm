package golang

// Permutations
// leetCode: https://leetcode-cn.com/problems/permutations/

// TODO 未完成
func Permute(nums []int) (ans [][]int) {
	path := make([]int, 0)
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(nums) {
			ans = append(ans, path)
		}
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}