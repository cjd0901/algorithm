package _022_06

// 每日温度
// leetcode: https://leetcode.cn/problems/iIQa4I/
func DailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	l := len(temperatures)
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			preIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIdx] = i - preIdx
		}
		stack = append(stack, i)
	}
	return ans
}
