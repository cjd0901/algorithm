package golang

// 204.Count Primes
// leetCode: https://leetcode-cn.com/problems/count-primes/

// 该算法会超时
func isPrime(n int) bool {
	if n == 3 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func CountPrimes(n int) int {
	if n == 0 || n == 1 || n == 2{
		return 0
	}
	count := 1
	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			count ++
		}
	}
	return count
}

// 官方题解：埃氏筛
// 如果一个数为质数，则这个数所有倍数的数为和数，在数组中置为1
// 为什么不从2*x开始标记而是从x*x开始标记？
// 因为2*x开始标记可能会被之前的数重复标记
func CountPrimes2(n int) int {
	list := make([]int, n)
	count := 0
	for i := 2; i < n; i++ {
		if list[i] == 0 {
			count += 1
			if i * i < n {
				for j := i * i; j < n; j += i {
					list[j] = 1
				}
			}
		}
	}
	return count
}