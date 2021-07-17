package _021_07

// 连续子数组的最大和
// leetcode: https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func MaximumSumOfConsecutiveSubarrays(nums []int) int {
	n, ans := len(nums), nums[0]
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1] + nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

func MaximumSumOfConsecutiveSubarrays2(nums []int) int {
	prev, ans := nums[0], nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		prev = max(prev + nums[i], nums[i])
		ans = max(prev, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}