package _022_02

// Maximum Difference Between Increasing Elements
// leetcode: https://leetcode-cn.com/problems/maximum-difference-between-increasing-elements/
func MaximumDifferenceBetweenIncreasingElements(nums []int) int {
	n := len(nums)
	ans := -1
	for i, min := 0, nums[0]; i < n; i++ {
		if nums[i] > min {
			ans = max(ans, nums[i] - min)
		} else {
			min = nums[i]
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}