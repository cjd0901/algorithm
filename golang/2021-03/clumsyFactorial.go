package _021_03

// Clumsy Factorial
// leetCode: https://leetcode-cn.com/problems/clumsy-factorial/
// clumsy(10) = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1

func ClumsyFactorial(N int) int {
	if N <= 2 {
		return N
	}else if N == 3 {
		return 6
	}else if N == 4 {
		return 7
	}
	mod := N % 4
	if mod == 0 {
		return N + 1
	}else if mod == 1 || mod == 2 {
		return N + 2
	}else {
		return N - 1
	}
}
