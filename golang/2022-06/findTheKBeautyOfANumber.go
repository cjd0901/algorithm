package _022_06

import "strconv"

// 找到一个数字的K美丽值
// leetcode: https://leetcode.cn/problems/find-the-k-beauty-of-a-number/
func FindTheKBeautyOfANumber(num int, k int) int {
	ans := 0
	numStr := strconv.Itoa(num)
	for i := 0; i+k <= len(numStr); i++ {
		n, _ := strconv.Atoi(numStr[i : i+k])
		if n == 0 {
			continue
		}
		if num%n == 0 {
			ans++
		}
	}
	return ans
}
