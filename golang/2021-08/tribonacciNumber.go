package _021_08

// N-th Tribonacci Number
// leetcode: https://leetcode-cn.com/problems/n-th-tribonacci-number/
func TribonacciNumber(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	q, w, e, r := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		q = w
		w = e
		e = r
		r = q + w + e
	}
	return r
}