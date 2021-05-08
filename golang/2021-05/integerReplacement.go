package _021_05

// Integer Replacement
// leetcode: https://leetcode-cn.com/problems/integer-replacement/
// Given a positive integer n, you can apply one of the following operations:
// If n is even, replace n with n / 2.
// If n is odd, replace n with either n + 1 or n - 1.
// Return the minimum number of operations needed for n to become 1.

// 位运算
func IntegerReplacement3(n int) int {
	count := 0
	for n != 1 {
		if n & 1 == 0 {
			n >>= 1
		}else {
			if n & 2 == 0 || n == 3 {
				n += -1
			}else {
				n++
			}
		}
		count++
	}
	return count
}

// dp：超时
func IntegerReplacement2(n int) int {
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		if isEven(i) {
			dp[i] = dp[i/2]+1
		}else if !isEven(i) && isEven((i+1)/2) {
			dp[i] = min(dp[(i+1)/2]+2, dp[i-1]+1)
		}else if !isEven(i) && isEven((i-1)/2) {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp[n]
}

func isEven(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}