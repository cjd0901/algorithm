package golang

// Permutations
// leetCode: https://leetcode-cn.com/problems/permutations/

func Permute(nums []int) (ans [][]int) {
	path := make([]int, 0)
	l := len(nums)
	used :=  make([]bool, l)
	var dfs func(index int)
	dfs = func(index int) {
		if index == l {
			temp := make([]int, l)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < l; i++ {
			if !used[i] {
				used[i] = true
				path = append(path, nums[i])
				//fmt.Println("回溯前==>", path, used, index)
				dfs(index+1)
				used[i] = false
				path = path[:len(path)-1]
				//fmt.Println("回溯后==>", path, used, index)
			}
		}
	}
	dfs(0)
	return
}