package _021_07

// Factorial Zeros
// leetcode: https://leetcode-cn.com/problems/factorial-zeros-lcci/
func FactorialZeros(n int) int {
	zeros := 0
	for n >= 5 {
		n = n/5
		zeros += n
	}
	return zeros
}