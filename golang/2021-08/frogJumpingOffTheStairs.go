package _021_08

// 青蛙跳台阶
// leetcode: https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func FrogJumpingOffTheStairs(n int) int {
	mod := int(1e9+7)
	if n == 0 {
		return 1
	}
	q, w, ans := 1, 1, 1
	for i := 2; i <= n; i++ {
		ans = (q + w)%mod
		q = w
		w = ans
	}
	return ans
}