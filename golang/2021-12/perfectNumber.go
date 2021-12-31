package _021_12

import "math"

// Perfect Number
// leetcode: https://leetcode-cn.com/problems/perfect-number/
func PerfectNumber(num int) bool {
	sum := 0
	sqrt := math.Sqrt(float64(num))
	for i := 1; i <= int(sqrt); i++ {
		if num % i == 0 {
			if i != num {
				sum += i
			}
			if num / i != num {
				sum += num / i
			}
		}
	}
	return sum == num
}