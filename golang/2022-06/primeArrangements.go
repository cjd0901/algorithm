package _022_06

// 质数排列
// leetcode: https://leetcode.cn/problems/prime-arrangements/
const mod = 1e9 + 7

func PrimeArrangements(n int) int {
	numsPrimes := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			numsPrimes++
		}
	}
	return factorial(numsPrimes) * factorial(n-numsPrimes) % mod
}

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func factorial(x int) int {
	f := 1
	for i := 1; i <= x; i++ {
		f = f * i % mod
	}
	return f
}
