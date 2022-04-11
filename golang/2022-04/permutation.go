package _022_04

// nowcoder: https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7?tpId=117&&tqId=37778&&companyId=239&rp=1&ru=/company/home/code/239&qru=/ta/job-code-high/question-ranking
func permutation(str string) []string {
	strBytes := []byte(str)
	l := len(strBytes)
	used := make([]bool, l)
	s := make([]byte, 0)
	ans := make([]string, 0)
	m := make(map[string]struct{})

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == l {
			if _, ok := m[string(s)]; !ok {
				ans = append(ans, string(s))
			}
			m[string(s)] = struct{}{}
			return
		}
		for i := 0; i < l; i++ {
			if !used[i] {
				used[i] = true
				s = append(s, strBytes[i])
				dfs(idx + 1)
				used[i] = false
				s = s[:len(s)-1]
			}
		}
	}
	dfs(0)
	return ans
}
