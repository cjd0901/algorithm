package _022_02

// Calculate Money in Leetcode Bank
// leetcode: https://leetcode-cn.com/problems/calculate-money-in-leetcode-bank/
func CalculateMoneyInLeetcodeBank(n int) int {
	sum := 0
	week, day := 0, 1  // 第一周起始为0，第一天起始为1
	for i := 0; i < n; i++ {
		money := week + day
		sum += money
		day++
		if day == 8 {
			day = 1
			week++
		}
	}
	return sum
}