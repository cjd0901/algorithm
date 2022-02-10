package _022_02

// Count Number of Pairs With Absolute Difference K
// leetcode: https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k/
func CountNumber(nums []int, k int) int {
	count := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if abs(nums[j] - nums[i]) == k {
				count++
			}
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}