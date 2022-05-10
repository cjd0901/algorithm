package _022_05

// leetcode: https://leetcode.cn/problems/di-string-match/
// 增减字符串匹配
func DiStringMatch(s string) []int {
	ans := make([]int, 0)
	l := len(s)
	min, max := 0, l
	for _, ch := range s {
		switch ch {
		case 'I':
			{
				ans = append(ans, min)
				min++
			}
		case 'D':
			{
				ans = append(ans, max)
				max--
			}
		}
	}
	ans = append(ans, min)
	return ans
}
